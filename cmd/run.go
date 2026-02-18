package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
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
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
