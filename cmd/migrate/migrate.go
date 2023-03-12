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

		if err := initDB(); err != nil {
			panic("init db instance failed!")
		}

		fmt.Println(global.DB.Config)
		// 3、迁移数据库
		modelMigrate()
	} else {
		genFile()
	}

}

func initDB() error {

	return database.Setup(config.DatabaseConfig.Driver)

}

func modelMigrate() {

}

func genFile() {

}
