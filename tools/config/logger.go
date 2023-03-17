package config

import "github.com/spf13/viper"

type Logger struct {
	Path       string
	Stdout     bool
	Level      string
	EnabledBus bool
	EnabledReq bool
	EnabledDb  bool
	EnabledJob bool
}

func InitLogger(v *viper.Viper) *Logger {
	return &Logger{
		Path:       v.GetString("path"),
		Stdout:     v.GetBool("sdtout"),
		Level:      v.GetString("level"),
		EnabledBus: v.GetBool("enabledbus"),
		EnabledReq: v.GetBool("enabledreq"),
		EnabledDb:  v.GetBool("enableddb"),
		EnabledJob: v.GetBool("enablejob"),
	}
}

var LoggerConfig = new(Logger)
