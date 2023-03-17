package config

import "github.com/spf13/viper"

type Gen struct {
	DBName    string
	FrontPath string
}

func InitGen(v *viper.Viper) *Gen {
	return &Gen{
		DBName:    v.GetString("dbname"),
		FrontPath: v.GetString("frontpath"),
	}
}

var GenConfig = new(Gen)
