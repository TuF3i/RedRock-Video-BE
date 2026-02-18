package cmd

import (
	"LiveDanmu/apps/gateway/danmu_gateway/server"

	"github.com/spf13/cobra"
)

// danmuCmd represents the danmu command
var danmuCmd = &cobra.Command{
	Use: "danmu",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunDanmuGateway()
	},
}

func init() {
	gatewayCmd.AddCommand(danmuCmd)
}
