package config

import (
	"encoding/json"
	"fly/tools"
	"fmt"

	"fly/tools/config"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "Get application config info",
	Example: "fly config -c config/settings.yml",
	Run: func(cmd *cobra.Command, args []string) {
		getConfigInfo()
	},
}

var configPath string

func init() {
	ConfigCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config/settings.yml", "Path to the configFile that program should read") // 给config命令添加flag，并设置了默认值
}
func getConfigInfo() {
	fmt.Println(tools.Yellow("now, begin to get the config info..."))

	/* jason-comment
	该函数内部，根据传入的配置文件路径，借助viper进行解析，解析后再将对应配置信息，解析到对应的结构体实例中，
	函数执行完毕，就可以通过config.xxConfig，访问内存中对应的配置段的信息了！
	*/
	config.Setup(configPath)

	// 接下来：开始做json的序列化输出，即完成config子命令的功能！

	application, err := json.MarshalIndent(config.ApplicationConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("application:", string(application))

	database, err := json.MarshalIndent(config.DatabaseConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("database:", string(database))

	jwt, err := json.MarshalIndent(config.JwtConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("jwt:", string(jwt))

	gen, err := json.MarshalIndent(config.GenConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("gen:", string(gen))

	logger, err := json.MarshalIndent(config.LoggerConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("logger:", string(logger))

	ssl, err := json.MarshalIndent(config.SslConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("ssl:", string(ssl))

	sync, err := json.MarshalIndent(config.SyncConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("sync:", string(sync))

	es, err := json.MarshalIndent(config.EsConfig, "", "	") // 返回的是[]byte 和 error
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("es:", string(es))
}
