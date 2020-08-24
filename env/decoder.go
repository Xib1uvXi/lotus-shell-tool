package env

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Decoder struct {
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) Decode(filePath string) *Config {
	var conf Config

	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		panic(fmt.Sprintf("decode toml file failed, err: %s", err.Error()))
	}

	return conf.validate()
}
