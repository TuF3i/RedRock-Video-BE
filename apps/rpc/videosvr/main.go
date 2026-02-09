package main

import (
	videosvr "LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr/videosvr"
	"log"
)

func main() {
	svr := videosvr.NewServer(new(VideoSvrImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
