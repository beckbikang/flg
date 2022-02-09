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
	defer lg.Sync()
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
	defer lg.Sync()
	lg.Info("a test")
	lg.Info("abc", zap.Int("int", 11))
}

func Benchmark_Write(b *testing.B) {
	var fconfig FConfig
	l := &Logger{}
	if _, err := toml.DecodeFile("./data/test.toml", &fconfig); err != nil {
		panic(err)
	}
	err := l.LoadFromObject(&fconfig)
	if err != nil {
		panic("TestLoadFromObject faild")
	}
	benchLog, err := l.GetLogByKey("test")
	for i := 1; i < b.N; i = i + 1 {
		benchLog.Info("abc", zap.Int("int", 11))
	}
}
