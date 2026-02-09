package main

import (
	"LiveDanmu/apps/rpc/videosvr/core/handle"
	videosvr "LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr/videosvr"
	"log"
)

func main() {
	svr := videosvr.NewServer(new(handle.VideoSvrImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
