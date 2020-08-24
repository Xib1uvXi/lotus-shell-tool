package env

import (
	"fmt"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	d := &Decoder{}

	conf := d.Decode("./test.toml")

	fmt.Println(conf)

	fmt.Println(conf.GetLogPath("ttt"))
}
