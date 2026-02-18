package cmd

import (
	"LiveDanmu/apps/gateway/user_gateway/server"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use: "user",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunUserGateway()
	},
}

func init() {
	gatewayCmd.AddCommand(userCmd)
}
