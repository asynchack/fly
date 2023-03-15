package config

import "github.com/spf13/viper"

type Sync struct {
	Cloud int
}

func InitSync(v *viper.Viper) *Sync {
	return &Sync{
		Cloud: v.GetInt("cloud"),
	}
}

var SyncConfig = new(Sync)
