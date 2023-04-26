package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	// o Select vai imprimir o canal que chagar a informação primerio;
	// Caso passe 4 segundo e não chagar a inforção em nenhum canal da timeout
	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)

	case msg2 := <-c2:
		fmt.Println("received", msg2)

	case <-time.After(time.Second * 4):
		fmt.Println("Timeout")
	}

}

type Message struct {
	ID  int64
	Msg string
}

func NewSelectTest() {
	channel1 := make(chan Message)
	channel2 := make(chan Message)
	var i int64
	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from RabbitMQ"}
			channel1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from Kafka"}
			channel2 <- msg
		}
	}()

	for {
		select {
		case msg := <-channel1: //RabbitMQ
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg.ID, msg.Msg)

		case msg := <-channel2: //RabbitMQ
			fmt.Printf("Received from kafka: ID: %d - %s\n", msg.ID, msg.Msg)
		}
	}
}
