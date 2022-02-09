# flg
a simple log lib wapper use zap and lumberjack



toml config 
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

how to use it with toml config  

1 create toml config file

2 import the lib,and then you can use it like this


common code 




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
	defer lt.Sync()
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



```golang

func TestLoadFromFile(t *testing.T){
	l := &Logger{}
	err := l.LoadFromFile("test.toml")
	if err != nil {
		panic("get file faild")
	}
	lg,err := l.GetLogByKey("test")
	lg.Sync()
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
	lg.Sync()
	lg.Info("a test")
	lg.Info("abc",zap.Int("int",11))
}
```

benchmark test

```shell
goos: darwin
goarch: arm64
pkg: flg
Benchmark_Write-8   	  334576	      3486 ns/op
PASS
ok  	flg	1.840s
```


