package main

import (
	"AssessMangerServer/app/service"
	_ "AssessMangerServer/boot"
	_ "AssessMangerServer/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	mq:=service.NewMQ()
	mq.Consume(service.QUEUE_NEWUDATA,"test1",service.SendMail)
	defer mq.Channel.Close()
	g.Server().Run()

}
