package messagequeue

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

// RabbitMQ Document https://www.tizi365.com/topic/18.html

func RunMQPublish() {
	// 连接RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建 RabbitMQ 信道
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// 声明需要操作的队列
	q, err := ch.QueueDeclare(
		"queue1", // 队列名
		false,    // 是否需要持久化
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	msgs, err := ch.Consume(
		q.Name, // 需要操作的队列名
		"",     // 消费者唯一id，不填，则自动生成一个唯一值
		true,   // 自动提交消息（即自动确认消息已经处理完成）
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	// 循环处理消息
	for d := range msgs {
		log.Printf("[消费者编号=%d] 收到消息: %s", 1, d.Body)
		// 模拟业务处理，休眠1秒
		time.Sleep(time.Second)
	}
}

func RunMQReceiver() {
	// 连接RabbitMQ Server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue1", // 队列名字
		false,    // 消息是否持久化
		false,    // 不使用的时候删除队列
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	// 消息内容
	body := "Hello World!"

	// 推送消息
	err = ch.Publish(
		"",     // exchange（交换机名字），这里忽略
		q.Name, // 路由参数，这里使用队列名字作为路由参数
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body), // 消息内容
		})
}

func RunMQ() {
	RunMQReceiver()
	RunMQPublish()
	
	time.Sleep(3 * time.Second)
}
