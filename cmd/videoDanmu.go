package cmd

import (
	"LiveDanmu/apps/consumer/video_danmu_consumer/server"

	"github.com/spf13/cobra"
)

// videoDanmuCmd represents the videoDanmu command
var videoDanmuCmd = &cobra.Command{
	Use: "videoDanmu",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunVideoDanmuConsumer()
	},
}

func init() {
	consumerCmd.AddCommand(videoDanmuCmd)
}
