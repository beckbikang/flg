# flg
a simple log lib wapper use zap and lumberjack

对zlog和lumberjack进行封装。希望能够结合两者的优点

toml config 基本的配置
```
[jackcfg]
filename="test.log"
maxsize=500
maxage=7
maxbackups=1000
localtime=true
compress=false


[zapcfgs]
[zapcfgs.1]
level="info"
isdev=true
logmod=3
servername="test"
```

how to use it with toml config  使用

1 create toml config file

2 import the lib,and then you can use it like this


common code 

一个典型的用法。你啥也不用关心，只用关心写啥日志就行。日志文件滚动啥的都交给日志文件。

你也可以定义多个日志类型的记录。

也支持LoadFromObject的方式创建日志对象


```golang

package main

import (
	"github.com/beckbikang/flg"
	"go.uber.org/zap"
)

//you can define you log var outer
var (
	lt *zap.Logger
	gflg *flg.Logger
)

func init(){
	gflg = &flg.Logger{}

	err := gflg.LoadFromFile("test.toml")
	if err != nil{
		panic("get file faild")
	}

	lt,err = gflg.GetLogByKey("test")
	if err != nil {
		panic(err)
	}
	lt.Info("start running")
}

func main() {

	// in my case you do not need define this
	//defer lt.Sync()

	lt.Info("a test")

}

```

测试例子

```golang

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

func TestLoadFromObject(t *testing.T){
	var fconfig FConfig
	l := &Logger{}
	if _, err := toml.DecodeFile("test.toml", &fconfig); err != nil {
		panic(err)
	}
	err := l.LoadFromObject(&fconfig)
	if err != nil {
		panic("TestLoadFromObject faild")
	}
	lg,err := l.GetLogByKey("test")
	lg.Info("a test")
	lg.Info("abc",zap.Int("int",11))
}
```


