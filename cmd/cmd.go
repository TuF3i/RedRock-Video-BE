package cmd

import (
	"fmt"
	"os"

	appDanmu "LiveDanmu/apps/consumer/live_danmu_consumer/server"
	appVideoDanmu "LiveDanmu/apps/consumer/video_danmu_consumer/server"
	dbInit "LiveDanmu/apps/db_init/auto-migrate/server"
	gwDanmu "LiveDanmu/apps/gateway/danmu_gateway/server"
	gwLive "LiveDanmu/apps/gateway/live_gateway/server"
	gwUser "LiveDanmu/apps/gateway/user_gateway/server"
	gwVideo "LiveDanmu/apps/gateway/video_gateway/server"
	rpcDanmu "LiveDanmu/apps/rpc/danmusvr/server"
	rpcLive "LiveDanmu/apps/rpc/livesvr/server"
	rpcUser "LiveDanmu/apps/rpc/usersvr/server"
	rpcVideo "LiveDanmu/apps/rpc/videosvr/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rv",
	Short: "RedRock-Video-BE Union Luncher",
	Long:  `Using this Luncher, you can run everything: gateway, consumer, rpc...`,
}

var dbInitCmd = &cobra.Command{
	Use:   "db-init",
	Short: "Usage: rv run db-init",
	Long:  `Example: rv run db-init`,
	Run: func(cmd *cobra.Command, args []string) {
		dbInit.RunInitDB()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Usage: rv run <unit> <name>",
	Long:  `Example: rv run rpc videosvr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: rv run <unit> <name>")
		fmt.Println("Example: rv run rpc videosvr")
		fmt.Println("----------")
		fmt.Printf("unit-gateway: ")
		fmt.Println("danmu | live | user | video")
		fmt.Printf("unit-rpc: ")
		fmt.Println("danmusvr | livesvr | usersvr | videosvr")
		fmt.Printf("consumer-rpc: ")
		fmt.Println("liveDanmu | videoDanmu ")
		fmt.Printf("init: ")
		fmt.Println("db-init ")
	},
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Usage: rv run <unit> <name>",
	Long:  `Example: rv run rpc videosvr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: rv run <unit> <name>")
		fmt.Println("Example: rv run rpc videosvr")
		fmt.Println("----------")
		fmt.Printf("unit-gateway: ")
		fmt.Println("danmu | live | user | video")
		fmt.Printf("unit-rpc: ")
		fmt.Println("danmusvr | livesvr | usersvr | videosvr")
		fmt.Printf("consumer-rpc: ")
		fmt.Println("liveDanmu | videoDanmu ")
	},
}

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Usage: rv run <unit> <name>",
	Long:  `Example: rv run rpc videosvr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: rv run <unit> <name>")
		fmt.Println("Example: rv run rpc videosvr")
		fmt.Println("----------")
		fmt.Printf("unit-gateway: ")
		fmt.Println("danmu | live | user | video")
		fmt.Printf("unit-rpc: ")
		fmt.Println("danmusvr | livesvr | usersvr | videosvr")
		fmt.Printf("consumer-rpc: ")
		fmt.Println("liveDanmu | videoDanmu ")
	},
}

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Usage: rv run <unit> <name>",
	Long:  `Example: rv run rpc videosvr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: rv run <unit> <name>")
		fmt.Println("Example: rv run rpc videosvr")
		fmt.Println("----------")
		fmt.Printf("unit-gateway: ")
		fmt.Println("danmu | live | user | video")
		fmt.Printf("unit-rpc: ")
		fmt.Println("danmusvr | livesvr | usersvr | videosvr")
		fmt.Printf("consumer-rpc: ")
		fmt.Println("liveDanmu | videoDanmu ")
	},
}

var danmuCmd = &cobra.Command{
	Use: "danmu",
	Run: func(cmd *cobra.Command, args []string) {
		gwDanmu.RunDanmuGateway()
	},
}

var liveCmd = &cobra.Command{
	Use: "live",
	Run: func(cmd *cobra.Command, args []string) {
		gwLive.RunLiveGateway()
	},
}

var userCmd = &cobra.Command{
	Use: "user",
	Run: func(cmd *cobra.Command, args []string) {
		gwUser.RunUserGateway()
	},
}

var videoCmd = &cobra.Command{
	Use: "video",
	Run: func(cmd *cobra.Command, args []string) {
		gwVideo.RunVideoGateway()
	},
}

var danmusvrCmd = &cobra.Command{
	Use: "danmusvr",
	Run: func(cmd *cobra.Command, args []string) {
		rpcDanmu.RunDanmuSvr()
	},
}

var livesvrCmd = &cobra.Command{
	Use: "livesvr",
	Run: func(cmd *cobra.Command, args []string) {
		rpcLive.RunLiveSvr()
	},
}

var usersvrCmd = &cobra.Command{
	Use: "usersvr",
	Run: func(cmd *cobra.Command, args []string) {
		rpcUser.RunUserSvr()
	},
}

var videosvrCmd = &cobra.Command{
	Use: "videosvr",
	Run: func(cmd *cobra.Command, args []string) {
		rpcVideo.RunVideoSvr()
	},
}

var liveDanmuCmd = &cobra.Command{
	Use: "liveDanmu",
	Run: func(cmd *cobra.Command, args []string) {
		appDanmu.RunLiveDanmuConsumer()
	},
}

var videoDanmuCmd = &cobra.Command{
	Use: "videoDanmu",
	Run: func(cmd *cobra.Command, args []string) {
		appVideoDanmu.RunVideoDanmuConsumer()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runCmd.AddCommand(gatewayCmd)
	runCmd.AddCommand(rpcCmd)
	runCmd.AddCommand(consumerCmd)
	runCmd.AddCommand(dbInitCmd)

	gatewayCmd.AddCommand(danmuCmd)
	gatewayCmd.AddCommand(liveCmd)
	gatewayCmd.AddCommand(userCmd)
	gatewayCmd.AddCommand(videoCmd)

	rpcCmd.AddCommand(danmusvrCmd)
	rpcCmd.AddCommand(livesvrCmd)
	rpcCmd.AddCommand(usersvrCmd)
	rpcCmd.AddCommand(videosvrCmd)

	consumerCmd.AddCommand(liveDanmuCmd)
	consumerCmd.AddCommand(videoDanmuCmd)
}
