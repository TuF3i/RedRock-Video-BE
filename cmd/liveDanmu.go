package cmd

import (
	"LiveDanmu/apps/consumer/live_danmu_consumer/server"

	"github.com/spf13/cobra"
)

// liveDanmuCmd represents the liveDanmu command
var liveDanmuCmd = &cobra.Command{
	Use: "liveDanmu",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunLiveDanmuConsumer()
	},
}

func init() {
	consumerCmd.AddCommand(liveDanmuCmd)
}
