package messagequeue

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	messageType string
	title       string
	content     string
}

type MessageQueue struct {
	messages   []Message
	subscribes []chan Message
	mutex      sync.Mutex
}

func (mq *MessageQueue) Publish(msg Message) {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	mq.messages = append(mq.messages, msg)

	for _, subscribe := range mq.subscribes {
		go func(ch chan Message) {
			ch <- msg
		}(subscribe)
	}
}

func (mq *MessageQueue) Subscribe() chan Message {
	ch := make(chan Message)
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	mq.subscribes = append(mq.subscribes, ch)
	return ch
}

func HandleConsumeMessage(ch chan Message) {
	for message := range ch {
		fmt.Println(message.title)
	}
}

func RunMessageQueue() {
	mq := MessageQueue{}

	ch1 := mq.Subscribe()
	ch2 := mq.Subscribe()

	go HandleConsumeMessage(ch1)
	go HandleConsumeMessage(ch2)

	//	Product data
	go func() {
		for i := 0; i < 10; i++ {
			msg := Message{messageType: "test", title: fmt.Sprintf("title%d", i), content: "test content"}
			mq.Publish(msg)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(12 * time.Second)

	close(ch1)
	close(ch2)
}
