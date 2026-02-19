package main

import (
	danmusvr "LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr/danmusvr"
	"log"
)

func main() {
	svr := danmusvr.NewServer(new(DanmuSvrImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
