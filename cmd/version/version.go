package version

import (
	"fly/tools"
	"fmt"

	"fly/common/global"

	"github.com/spf13/cobra"
)

// 该子命令输出版本号信息

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "print the version number for fly",
	Long:    "print the version number for fly, same as short desc, haha!",
	Example: "fly version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tools.Yellow("fly version: " + global.Version))

	},
}
