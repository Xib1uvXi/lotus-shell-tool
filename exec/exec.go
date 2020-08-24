package exec

import (
	bind_data "github.com/Xib1uvXi/lotus-shell-tool/bind-data"
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

func (e Executor) StartLotus() error {
	checkCmdExist("lotus")
	name := "lotus-blockchain"
	shB, err := bind_data.Asset("scripts/start_lotus.sh")
	if err != nil {
		log.Error("exec start lotus shell failed", "msg: ", err)
		return err
	}

	if err := execCmdByTmpFile(shB, e.conf.Env(), name, e.conf.GetLogPath(name)); err != nil {
		return err
	}

	return nil
}

//func (e Executor) StartMiner() error {
//
//}
//
//func (e Executor) StopMiner() error {
//
//}
//
//func (e Executor) StartWorker() error {
//
//}
//
//func (e Executor) StopWorker() error {
//
//}

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
			log.Debug("exec cmd failed", "result", string(rb))
			return err
		}
	}
	return
}

func checkCmdExist(cmd string) {
	rb, _ := exec.Command("which", cmd).Output()
	if !strings.Contains(string(rb), cmd) {
		panic(cmd + " not installed")
	}
}
