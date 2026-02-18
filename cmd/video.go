package cmd

import (
	"LiveDanmu/apps/gateway/video_gateway/server"

	"github.com/spf13/cobra"
)

// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use: "video",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunVideoGateway()
	},
}

func init() {
	gatewayCmd.AddCommand(videoCmd)
}
