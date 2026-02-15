package main

import (
	livesvr "LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr/livesvr"
	"log"
)

func main() {
	svr := livesvr.NewServer(new(LiveSvrImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
