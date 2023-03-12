package config

import "github.com/spf13/viper"

type Jwt struct {
	Secret  string
	Timeout int64
}

func InitJwt(v *viper.Viper) *Jwt {
	return &Jwt{
		Secret:  v.GetString("secret"),
		Timeout: v.GetInt64("timeout"),
	}
}

var JwtConfig = new(Jwt)
