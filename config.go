package flg




const (
	//logmod 默认是 1 文件  2 stdout 4 其他
	FILE_MODE = 1<<iota
	STDOUT_MODE
	OTHER_MODE

	TIME_KEY = "time"
	LEVEL_KEY = "level"
	NAME_KEY = "logger"
	CALLER_KEY = "line"
	MESSAGE_KEY = "data"
	STACK_TRACE_KEY = "stacktrace"
	SERVER_NAME = "server"

)




type FConfig struct {
	Lfg LConfig `toml:"jackcfg"`

	Zfgs map[string]ZConfig`toml:"zapcfgs"`
}


type LConfig struct{
	Filename string `toml:"filename"`
	MaxSize int `toml:"maxsize"` //MB
	MaxAge int `toml:"maxage"`
	MaxBackups int `toml:"maxbackups"`
	LocalTime bool `toml:"localtime"`
	Compress bool `toml:"compress"`
}


type ZConfig struct{
	LogName string `toml:"logname"`
	Timekey string `toml:"timekey"`
	LevelKey string `toml:"levelkey"`
	NameKey string `toml:"namekey"`
	CallerKey string `toml:"callerkey"`
	MessageKey string `toml:"messagekey"`
	StacktraceKey string `toml:"stacktracekey"`
	
	Level string  `toml:"level"`
	IsDev bool `toml:"isdev"`
	LogMod int8 `toml:"logmod"`
	ServerName string `toml:"servername"`
}

