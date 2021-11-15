package service

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/streadway/amqp"
	"log"
)

var MQConn *amqp.Connection
func init()  {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", g.Cfg().Get("RabbitMq.user"), g.Cfg().Get("RabbitMq.pwd"), g.Cfg().Get("RabbitMq.server"), g.Cfg().Get("RabbitMq.port"))
	conn, err := amqp.Dial(dsn)
	if err!=nil{
		log.Fatal(err)
	}
	MQConn=conn

	log.Print(MQConn.Major)
}
func GetConn() *amqp.Connection {
	return MQConn
}
