package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/* jason-comment
借助viper库实现，根据给的的配置文件路径，加载、并解析其中配置信息，映射到本包中，其他文件中定义的结构体中，（实现配置文件->程序内存中变量的转换）
*/

func Setup(filePath string) {
	/* jason-comment
	为什么要用？ viper的sub方法，得到*viper.Viper实例，然后再调用一个结构体方法，把该实例，转为一个结构体实例，再返回？
	原因如下：没理解？
	ref: https://github.com/spf13/viper#extracting-a-sub-tree

	*/
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error occur when read config file, %q\n", err)

	}
	// fmt.Println(viper.AllSettings())

	// app配置初始化
	application := viper.Sub("settings.application")
	if application == nil {
		panic("file of configuration not found!")

	}

	ApplicationConfig = InitApplication(application) // 赋值给：包导入时的，初始化的一个空的ApplicationConfig使得它有值了

	// logger配置初始化
	logger := viper.Sub("settings.logger")
	if logger == nil {
		panic("logger configuration not found!")
	}

	LoggerConfig = InitLogger(logger)

	// db配置初始化

	database := viper.Sub("settings.database")
	if database == nil {
		panic("database configuration not found!")
	}
	DatabaseConfig = InitDatabase(database)

	// jwt配置初始化

	jwt := viper.Sub("settings.jwt")

	if jwt == nil {
		panic("jwt configuration not found!")
	}
	JwtConfig = InitJwt(jwt)

	// sync配置初始化

	sync := viper.Sub("settings.sync")
	if sync == nil {
		panic("sync configuration not found!")
	}
	SyncConfig = InitSync(sync)

	// gen配置初始化

	gen := viper.Sub("settings.gen")
	if gen == nil {
		panic("gen configuration not found!")
	}
	GenConfig = InitGen(gen)

	// ssl配置初始化
	ssl := viper.Sub("settings.ssl")
	if ssl == nil {
		panic("ssl configuration not found!")
	}

	SslConfig = InitSsl(ssl)

	// es配置初始化，暂略

	es := viper.Sub("settings.es")
	if es == nil {
		panic("es configuration not found!")
	}
	EsConfig = InitEs(es)

	// fmt.Println(ApplicationConfig, LoggerConfig, SslConfig, GenConfig, SyncConfig, JwtConfig, DatabaseConfig)

}
