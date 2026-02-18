package cmd

import (
	app "LiveDanmu/apps/rpc/livesvr/server"

	"github.com/spf13/cobra"
)

// livesvrCmd represents the livesvr command
var livesvrCmd = &cobra.Command{
	Use: "livesvr",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunLiveSvr()
	},
}

func init() {
	rpcCmd.AddCommand(livesvrCmd)
}
