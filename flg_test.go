package flg

import (
	"testing"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

var (
	lg *zap.Logger
)

/**
测试初始化
*/
func TestLoadFromFile(t *testing.T) {
	l := &Logger{}
	err := l.LoadFromFile("./data/test.toml")
	if err != nil {
		panic("get file faild")
	}
	lg, err := l.GetLogByKey("test")
	lg.Info("a test")

	lg.Info("abc", zap.Int("int", 11))
}

func TestLoadFromObject(t *testing.T) {
	var fconfig FConfig
	l := &Logger{}
	if _, err := toml.DecodeFile("./data/test.toml", &fconfig); err != nil {
		panic(err)
	}
	err := l.LoadFromObject(&fconfig)
	if err != nil {
		panic("TestLoadFromObject faild")
	}
	lg, err := l.GetLogByKey("test")
	lg.Info("a test")
	lg.Info("abc", zap.Int("int", 11))
}
