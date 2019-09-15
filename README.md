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

```
package main

import "github.com/beckbikang/flg"

func main() {

	l := &flg.Logger{}

	err := l.LoadFromFile("test.toml")
	if err != nil{
		panic("get file faild")
	}
	ltest,err := l.GetLogByKey("test")
	if err != nil {
		panic(err)
	}
	defer ltest.Sync()

	ltest.Info("a test")

}
```


