package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// rabbitmq
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from rabbitmq"}
			c1 <- msg
		}
	}()

	// kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from kafka"}
			c2 <- msg
		}
	}()

	for { // loop infinito para ficar recebendo mensagens o tempo todo
		select {
		case msg := <-c1: // rabbitmq
			fmt.Printf("Received from RabbiMQ: %d - %s\n", msg.Id, msg.Msg)

		case msg := <-c2: // kafka
			fmt.Printf("Received from Kafka: %d - %s\n", msg.Id, msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
