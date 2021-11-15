package service

import (
	"github.com/streadway/amqp"
	"log"
	"strings"
)

const (
	QUEUE_NEWUDATA="test1"  //数据新增成功
	QUEUE_DATAMODIFY="test2"  //数据修改成功
	QUEUE_ASSESS="test3" //发送评估

)

type MQ struct {
	Channel *amqp.Channel
}
func NewMQ() *MQ  {
	c,err:=GetConn().Channel()
	if err!=nil{
		log.Println(err)
		return nil
	}

	return &MQ{Channel:c}
}

//申明队列以及绑定路由key
//多个队列 可以用逗号分隔
func(this *MQ) DecQueueAndBind (queues string,key string,exchange string) error{
	qList:=strings.Split(queues,",")
	for _,queue:=range qList{
		q,err:=this.Channel.QueueDeclare(queue,false,false,false,false,nil)
		if err!=nil{
			return err
		}
		err=this.Channel.QueueBind(q.Name,key,exchange,false,nil)
		if err!=nil{
			return err
		}
	}
	return  nil
}

func(this *MQ) Consume(queue string,key string,callbak func(<-chan amqp.Delivery,string) ){
	msgs,err:=this.Channel.Consume(queue,key,false,false,false,false,nil)
	if err!=nil{
		log.Fatal(err)
	}
	callbak(msgs,key)
}
