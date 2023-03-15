package config

import "github.com/spf13/viper"

type Ssl struct {
	Domain string
	Enable string
	Key    string
	Pem    string
}

func InitSsl(v *viper.Viper) *Ssl {
	return &Ssl{
		Domain: v.GetString("domain"),
		Enable: v.GetString("enable"),
		Key:    v.GetString("key"),
		Pem:    v.GetString("pem"),
	}
}

var SslConfig = new(Ssl)
