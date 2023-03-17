package server

import (
	"fly/tools"
	"fmt"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Short:   "start the cmdb backend！",
	Example: "fly server -c config/settings.yml",
	Run: func(cmd *cobra.Command, args []string) {
		startApp()
	},
}

func startApp() {
	fmt.Println(tools.Yellow("now, begin to start the App..."))
}
