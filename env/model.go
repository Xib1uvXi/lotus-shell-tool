package env

import (
	"fmt"
	"path/filepath"
	"time"
)

const (
	_521MiB = "512MiB"
	_32GiB  = "32GiB"
)

type Config struct {
	Name   string `toml:"NAME"`
	Size   string `toml:"SIZE"`
	Proofs *Proofs
	Path   *Path
	Worker *Worker
}

func (c *Config) GetLogPath(id string) string {
	return filepath.Join(c.Path.LogPath, fmt.Sprintf("%s-%s-%s-%s.log", id, c.Name, c.Size, time.Now().Format("20060102150405")))
}

func (c *Config) Env() []string {
	return []string{
		fmt.Sprintf("IPFS_GATEWAY=%s", c.Proofs.IpfsGateway),
		fmt.Sprintf("FIL_PROOFS_USE_GPU_COLUMN_BUILDER=%s", c.Proofs.FilProofsUseGpuColumnBuilder),
		fmt.Sprintf("FIL_PROOFS_USE_GPU_TREE_BUILDER=%s", c.Proofs.FilProofsUseGpuTreeBuilder),
		fmt.Sprintf("FIL_PROOFS_PARAMETER_CACHE=%s", c.Proofs.FilProofsMaximizeCaching),
		fmt.Sprintf("FIL_PROOFS_PARAMETER_CACHE=%s", c.Path.FilProofsParameterCache),
		fmt.Sprintf("TMPDIR=%s", c.Path.TmpDir),
		fmt.Sprintf("LOTUS_PATH=%s", c.Path.LotusPath),
		fmt.Sprintf("LOTUS_STORAGE_PATH=%s", c.Path.LotusStoragePath),
		fmt.Sprintf("LOG_PATH=%s", c.Path.LogPath),
		fmt.Sprintf("WORKER_REPO=%s", c.Path.WorkerRepo),
		fmt.Sprintf("RUST_LOG=%s", c.Path.RustLog),
		fmt.Sprintf("MINER_API_INFO=%s", c.Path.MinerApiInfo),
	}
}

type Proofs struct {
	IpfsGateway string `toml:"IPFS_GATEWAY"`
	// todo https://github.com/filecoin-project/rust-fil-proofs/#advanced-gpu-usage
	FilProofsUseGpuColumnBuilder string `toml:"FIL_PROOFS_USE_GPU_COLUMN_BUILDER"`
	FilProofsUseGpuTreeBuilder   string `toml:"FIL_PROOFS_USE_GPU_TREE_BUILDER"`
	FilProofsMaximizeCaching     string `toml:"FIL_PROOFS_PARAMETER_CACHE"`
}

type Path struct {
	FilProofsParameterCache string `toml:"FIL_PROOFS_PARAMETER_CACHE"`
	TmpDir                  string `toml:"TMPDIR"`
	LotusPath               string `toml:"LOTUS_PATH"`
	LotusStoragePath        string `toml:"LOTUS_STORAGE_PATH"`
	LogPath                 string `toml:"LOG_PATH"`
	WorkerRepo              string `toml:"WORKER_REPO"`
	RustLog                 string `toml:"RUST_LOG"`
	MinerApiInfo            string `toml:"MINER_API_INFO"`
}

type Worker struct {
	Listen string `toml:"LISTEN"`
	AP     bool   `toml:"AP"`
	P1     bool   `toml:"P1"`
	P2     bool   `toml:"P2"`
	C      bool   `toml:"C"`
}

func (c *Config) validate() *Config {
	if c.Name == "" {
		c.Name = "default"
	}

	if c.Size == "" {
		panic("toml config need set SIZE")
	}

	if c.Size != _32GiB && c.Size != _521MiB {
		panic(fmt.Sprintf("SIZE only support %s or %s, you set %s", _521MiB, _32GiB, c.Size))
	}

	if c.Proofs.IpfsGateway == "" {
		c.Proofs.IpfsGateway = "https://proof-parameters.s3.cn-south-1.jdcloud-oss.com/ipfs/"
	}

	if c.Proofs.FilProofsUseGpuColumnBuilder == "" {
		c.Proofs.FilProofsUseGpuColumnBuilder = "1"
	}

	if c.Proofs.FilProofsUseGpuTreeBuilder == "" {
		c.Proofs.FilProofsUseGpuTreeBuilder = "1"
	}

	if c.Proofs.FilProofsMaximizeCaching == "" {
		c.Proofs.FilProofsMaximizeCaching = "1"
	}

	if c.Path.FilProofsParameterCache == "" {
		panic("toml config need set FIL_PROOFS_PARAMETER_CACHE")
	}

	if c.Path.TmpDir == "" {
		panic("toml config need set TMPDIR")
	}

	//if c.Path.LotusPath == "" {
	//	panic("toml config need set LOTUS_PATH")
	//}
	//
	//if c.Path.LotusStoragePath == "" {
	//	panic("toml config need set LOTUS_STORAGE_PATH")
	//}

	if c.Path.LogPath == "" {
		panic("toml config need set LOG_PATH")
	}

	if c.Path.RustLog == "" {
		c.Path.RustLog = "info"
	}

	return c
}
