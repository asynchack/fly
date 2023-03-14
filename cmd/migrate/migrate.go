package migrate

import (
	"fly/common/database"
	"fly/common/global"
	"fly/tools"
	"fly/tools/config"
	"fmt"

	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "migrate and init the table!",
	Example: "fly migrate -c config/settings.yml",
	Run: func(cmd *cobra.Command, args []string) {
		initDatabase()
	},
}

var (
	configFilePath string
	generate       bool
	goAdmin        bool
)

func init() {
	MigrateCmd.PersistentFlags().StringVarP(&configFilePath, "cofig", "c", "config/settings.yml", "init table ")
	MigrateCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate the file which for init table, you must generate at least one time before migrate")
	MigrateCmd.PersistentFlags().BoolVarP(&goAdmin, "goAdmin", "", false, "goadmin (ignore for now)")
}

func initDatabase() {
	fmt.Println(tools.Yellow("now, begin to init the database..."))

	if !generate {
		// 1、加载配置
		config.Setup(configFilePath)

		// 2、实例化数据库连接 globa.DB

		initDB()

		fmt.Println(global.DB.Config)
		// 3、迁移数据库
		modelMigrate()
	} else {
		genFile()
	}

}

func initDB() {

	if err := database.Setup(config.DatabaseConfig.Driver); err != nil {
		panic("init db instance failed!")
	}
	// 此时成功的话：globa.DB已经指向了全局的*gorm.DB实例，可以做一些全局配置
	if config.DatabaseConfig.Driver == "mysql" {
		global.DB.Set("gorm:table_options", "ENGINE=InonoDB CHARSET=utf8mb4")
	}

	fmt.Println(tools.Yellow("开始migrate所有模型"))

	modelMigrate()

	fmt.Println(tools.Yellow("migreate所有表完成"))

}

func modelMigrate() {
	// 先前一个Migration表，该表记录，每个表（模型）是否迁移，以及迁移apply的时间是多少
	// 首先利用Debug（）方法，获得一个开启了debug日志级别的Session，Session也是*gorm.DB的类型（是它的子集？） 2者关系？
	// 在这个session中进行migrate，应该是仅仅该session上迁移的表，获得更详细的日志 ，日志输出？
	global.DB.Debug().AutoMigrate()
}

func genFile() {

}
