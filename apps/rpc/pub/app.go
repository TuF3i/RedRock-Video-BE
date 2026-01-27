package pub

import (
	pub "LiveDanmu/apps/rpc/pub/kitex_gen/pub/pub"
	"log"
)

func main() {
	svr := pub.NewServer(new(PubImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
