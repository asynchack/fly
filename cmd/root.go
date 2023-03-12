package cmd

import (
	"errors"
	"fmt"
	"os"

	"fly/tools"

	"fly/cmd/config"
	"fly/cmd/migrate"
	"fly/cmd/server"
	"fly/cmd/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fly",
	Short: "fly is a CMDB backend server",
	Long: `fly is a CMDB backend server,	So, let us fly`,

	Args: func(cmd *cobra.Command, args []string) error {
		// 用于处理 命令行参数的回调函数，可以判断参数数量
		if len(args) < 1 {
			tip()
			return errors.New(tools.Red("At least one args to Use!"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 单独执行主程序，无任何参数和command情况时：
		// Run是为一个command定义执行动作的地方，最顶层的fly没有实际动作，只输出提示
		tip()
	},
}

func tip() {
	usageString := "欢迎使用 fly CMDB系统!" + tools.Yellow("fly -h") + "查看使用帮助"
	fmt.Printf("%s\n", usageString)
}
func init() {
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(migrate.MigrateCmd)
	rootCmd.AddCommand(server.ServerCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
