package flg

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	TIMEFORMAT = "2006-01-02 15-04-05.000"
)

//base struct
type Logger struct {
	zlogs   map[string]*zap.Logger
	once    sync.Once
	fconfig *FConfig
}

//format
func currentTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	t = t.Local()
	enc.AppendString(t.Format(TIMEFORMAT))
}

//get data
func (l *Logger) GetLogByKey(key string) (*zap.Logger, error) {
	zlog, ok := l.zlogs[key]
	if !ok {
		return nil, errors.New("logger is not exist")
	}
	return zlog, nil
}

//get logger from object
func (l *Logger) LoadFromObject(fconfig *FConfig) error {
	l.once.Do(
		func() {
			l.fconfig = fconfig
			l.zlogs = make(map[string]*zap.Logger)
			err := l.makeLogger()
			if err != nil {
				panic("make logger faild")
			}
		})
	if l.fconfig == nil || len(l.zlogs) == 0 {
		return errors.New("log config not exist or make log faild")
	}
	return nil
}

//get log from toml file
func (l *Logger) LoadFromFile(filename string) error {
	//check file
	_, e := os.Stat(filename)
	if e != nil {
		return errors.New("file not exist")
	}
	l.once.Do(
		func() {
			var fconfig FConfig
			if _, err := toml.DecodeFile(filename, &fconfig); err != nil {
				panic(err)
			}
			l.fconfig = &fconfig
			l.zlogs = make(map[string]*zap.Logger)
			err := l.makeLogger()
			if err != nil {
				panic("make logger faild")
			}
		})

	if l.fconfig == nil || len(l.zlogs) == 0 {
		return errors.New("log config not exist or make log faild")
	}
	return nil
}

//set default config
func setDefaultConfig(zconfig *ZConfig) {
	if len(zconfig.Timekey) == 0 {
		zconfig.Timekey = TIME_KEY
	}

	if len(zconfig.LevelKey) == 0 {
		zconfig.LevelKey = LEVEL_KEY
	}

	if len(zconfig.NameKey) == 0 {
		zconfig.NameKey = NAME_KEY
	}
	if len(zconfig.CallerKey) == 0 {
		zconfig.CallerKey = CALLER_KEY
	}
	if len(zconfig.MessageKey) == 0 {
		zconfig.MessageKey = MESSAGE_KEY
	}
	if len(zconfig.StacktraceKey) == 0 {
		zconfig.StacktraceKey = STACK_TRACE_KEY
	}

}

//build logger
func (l *Logger) makeLogger() error {
	zfgs := l.fconfig.Zfgs

	lj := l.makeLumberjackLog()
	//创建可用的日志对象
	for _, zfg := range zfgs {
		setDefaultConfig(&zfg)
		encoderConfig := l.makeEncoderConfig(&zfg)

		// 设置日志级别
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.UnmarshalText([]byte(zfg.Level))

		zws := make([]zapcore.WriteSyncer, 0)

		zws = append(zws, zapcore.AddSync(&lj))

		if 1 == (zfg.LogMod & STDOUT_MODE) {
			zws = append(zws, zapcore.AddSync(os.Stdout))
		}

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zws...),
			atomicLevel,
		)

		zapOptions := make([]zap.Option, 0)

		//get aller
		caller := zap.AddCaller()
		zapOptions = append(zapOptions, caller)

		if zfg.IsDev {
			//add file and line
			development := zap.Development()
			zapOptions = append(zapOptions, development)
		}

		filed := zap.Fields(zap.String(SERVER_NAME, zfg.ServerName))
		zapOptions = append(zapOptions, filed)

		l.zlogs[zfg.ServerName] = zap.New(core, zapOptions...)
	}
	return nil
}

//make zconfig
func (l *Logger) makeEncoderConfig(zconfig *ZConfig) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        zconfig.Timekey,
		LevelKey:       zconfig.LevelKey,
		NameKey:        zconfig.NameKey,
		CallerKey:      zconfig.CallerKey,
		MessageKey:     zconfig.MessageKey,
		StacktraceKey:  zconfig.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     currentTimeEncoder,             // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
}

//get lumber jacklog
func (l *Logger) makeLumberjackLog() lumberjack.Logger {
	lkcfg := l.fconfig.Lfg
	lj := lumberjack.Logger{
		Filename:   lkcfg.Filename,   // 日志文件路径
		MaxSize:    lkcfg.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: lkcfg.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     lkcfg.MaxAge,     // 文件最多保存多少天
		Compress:   lkcfg.Compress,   //压缩
	}
	return lj
}
