package config

import "github.com/spf13/viper"

type Es struct {
	Enabled  bool
	Host     string
	User     string
	Password string
	Index    string
}

func InitEs(v *viper.Viper) *Es {
	return &Es{
		Enabled:  v.GetBool("enabled"),
		Host:     v.GetString("host"),
		User:     v.GetString("user"),
		Password: v.GetString("password"),
		Index:    v.GetString("index"),
	}
}

var EsConfig = new(Es)
