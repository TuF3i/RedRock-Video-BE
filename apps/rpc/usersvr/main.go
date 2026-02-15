package main

import (
	usersvr "LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr/usersvr"
	"log"
)

func main() {
	svr := usersvr.NewServer(new(UserSvrImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
