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
	if err != nil{
		panic("get file faild")
	}
	ltest,err := l.GetLogByKey("test")
	ltest.Info("aaaaa")

}