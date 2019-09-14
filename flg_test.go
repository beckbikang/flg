package flg

import (
	"testing"
	"go.uber.org/zap"
)


var(

	lg *zap.Logger
)

/**
 测试初始化
 */
func TestLoadFromFile(t *testing.T){

	l := &Logger{}
	err := l.LoadFromFile("test.toml")
	if err != nil {
		panic("get file faild")
	}
	lg,err := l.GetLogByKey("test")
	lg.Info("a test")

	lg.Info("abc",zap.Int("int",11))



}