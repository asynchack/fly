package config

import "github.com/spf13/viper"

type Database struct {
	Driver string
	Source string
}

func InitDatabase(v *viper.Viper) *Database {
	return &Database{
		Driver: v.GetString("driver"),
		Source: v.GetString("source"),
	}

}

var DatabaseConfig = new(Database)
