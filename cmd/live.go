package cmd

import (
	"LiveDanmu/apps/gateway/live_gateway/server"

	"github.com/spf13/cobra"
)

// liveCmd represents the live command
var liveCmd = &cobra.Command{
	Use: "live",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunLiveGateway()
	},
}

func init() {
	gatewayCmd.AddCommand(liveCmd)
}
