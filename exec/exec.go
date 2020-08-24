package exec

import (
	"fmt"
	"github.com/Xib1uvXi/lotus-shell-tool/env"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var log *zap.SugaredLogger

func init() {
	tmpLog, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	log = tmpLog.Sugar()
}

type Executor struct {
	conf *env.Config
}

func NewExecutor(conf *env.Config) *Executor {
	return &Executor{conf: conf}
}

func (e *Executor) StartLotus() error {
	checkCmdExist("lotus")
	name := "lotus-blockchain"

	if e.conf.Path.LotusPath == "" {
		panic("config toml need set LOTUS_PATH")
	}

	execCmd := fmt.Sprintf("export PATH=$PATH:~/tools/filecoin/calibration; exec -a %v lotus daemon >>%v 2>&1 &", name, e.conf.GetLogPath(name))

	if err := execCmdByTmpFile([]byte(execCmd), e.conf.Env()); err != nil {
		return err
	}

	return nil
}

func (e *Executor) StartMiner() error {
	checkCmdExist("lotus-miner")
	name := e.conf.Name

	if e.conf.Path.LotusStoragePath == "" {
		panic("config toml need set LOTUS_STORAGE_PATH")
	}

	if e.conf.Path.MinerApiInfo == "" {
		panic("config toml need set MINER_API_INFO")
	}

	execCmd := fmt.Sprintf("export PATH=$PATH:~/tools/filecoin/calibration; exec -a %v nohup lotus-miner run >>%v 2>&1 &", name, e.conf.GetLogPath(name))

	if err := execCmdByTmpFile([]byte(execCmd), e.conf.Env()); err != nil {
		return err
	}

	return nil
}

func (e *Executor) Stop(force bool) error {
	name := e.conf.Name
	return KillLocalProcess(name, force)
}

func (e Executor) StartWorker() error {
	checkCmdExist("lotus-miner")
	name := e.conf.Name

	if e.conf.Path.WorkerRepo == "" {
		panic("config toml need set WORKER_REPO")
	}

	if e.conf.Path.MinerApiInfo == "" {
		panic("config toml need set MINER_API_INFO")
	}

	if e.conf.Path.MinerApiInfo == "" {
		panic("config toml need set MINER_API_INFO")
	}

	execCmd := fmt.Sprintf(
		"mkdir -p %v;export PATH=$PATH:~/tools/filecoin/calibration; exec -a %v nohup lotus-worker --worker-repo=%v run --listen=%v --addpiece=%v --precommit1=%v --precommit2=%v --commit=%v >>%v 2>&1 &",
		e.conf.Path.WorkerRepo, name, e.conf.Path.WorkerRepo, e.conf.Worker.Listen, e.conf.Worker.AP, e.conf.Worker.P1, e.conf.Worker.P2, e.conf.Worker.C, e.conf.GetLogPath(name))

	if err := execCmdByTmpFile([]byte(execCmd), e.conf.Env()); err != nil {
		return err
	}

	return nil
}

func execCmdByTmpFile(cmd []byte, env []string, args ...string) (err error) {
	tmpFilePath := filepath.Join("/tmp", bson.NewObjectId().Hex()+".sh")

	if err = ioutil.WriteFile(tmpFilePath, cmd, 0755); err != nil {
		return
	}
	defer os.Remove(tmpFilePath)

	log.Info("exec cmd by file", "sh file", tmpFilePath, "args", args, "args len", len(args), "cmd content", string(cmd))

	args = append([]string{tmpFilePath}, args...)

	runCmd := exec.Command("/bin/bash", args...)
	runCmd.Env = env

	if rb, err := runCmd.Output(); err != nil {
		if len(rb) != 0 {
			log.Error("exec cmd failed", "result", string(rb))
			return err
		}
	}
	return
}

// 杀死某个本地程序
func KillLocalProcess(appName string, force bool) (err error) {
	_, err = exec.Command("/bin/bash", "-c", killCmd(appName, force)).Output()
	return
}

func killCmd(appName string, force bool) string {
	// kill -9无法被signal chan收到
	if force {
		return fmt.Sprintf("ps aux|grep \"%v\"|awk '{print $2}'|xargs kill -9", appName)
	}
	return fmt.Sprintf("ps aux|grep \"%v\"|awk '{print $2}'|xargs kill -2", appName)
}

func checkCmdExist(cmd string) {
	rb, _ := exec.Command("which", cmd).Output()
	if !strings.Contains(string(rb), cmd) {
		panic(cmd + " not installed")
	}
}
