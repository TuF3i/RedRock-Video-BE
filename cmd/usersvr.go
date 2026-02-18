package cmd

import (
	app "LiveDanmu/apps/rpc/usersvr/server"

	"github.com/spf13/cobra"
)

// usersvrCmd represents the usersvr command
var usersvrCmd = &cobra.Command{
	Use: "usersvr",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunUserSvr()
	},
}

func init() {
	rpcCmd.AddCommand(usersvrCmd)
}
