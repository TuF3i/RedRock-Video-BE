package cmd

import (
	app "LiveDanmu/apps/rpc/danmusvr/server"

	"github.com/spf13/cobra"
)

// danmusvrCmd represents the danmusvr command
var danmusvrCmd = &cobra.Command{
	Use: "danmusvr",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunDanmuSvr()
	},
}

func init() {
	rpcCmd.AddCommand(danmusvrCmd)
}
