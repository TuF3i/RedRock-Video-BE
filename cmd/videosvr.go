package cmd

import (
	app "LiveDanmu/apps/rpc/videosvr/server"

	"github.com/spf13/cobra"
)

// videosvrCmd represents the videosvr command
var videosvrCmd = &cobra.Command{
	Use: "videosvr",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunVideoSvr()
	},
}

func init() {
	rpcCmd.AddCommand(videosvrCmd)
}
