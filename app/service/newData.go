package service

import (
	"fmt"
	"github.com/streadway/amqp"
)

func SendMail(msgs <-chan amqp.Delivery ,c string )  {
	for msg:=range msgs{
		fmt.Println("收到消息",string(msg.Body))
		msg.Ack(true)
	}

}
