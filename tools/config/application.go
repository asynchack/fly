package config

import "github.com/spf13/viper"

type Application struct {
	Mode          string
	Host          string
	Name          string
	Port          string
	Readtimeout   int
	Writertimeout int
	Enableldap    bool
}

func InitApplication(v *viper.Viper) *Application {
	return &Application{
		Mode:          v.GetString("mode"),
		Host:          v.GetString("host"),
		Name:          v.GetString("name"),
		Port:          v.GetString("port"),
		Readtimeout:   v.GetInt("readtimeout"),
		Writertimeout: v.GetInt("writertimeout"),
		Enableldap:    v.GetBool("enableldap"),
	}
}

var ApplicationConfig = new(Application) // 包被导入时，初始化一个空的Application实例，为nil ；当执行了config.Setup之后，会将读取的配置文件内容，填充到该结构体，进行赋值
// 且是大写，之后就可以被，包外其他包获取，从而根据配置文件中的值，进行逻辑决策！
