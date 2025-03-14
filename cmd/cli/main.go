package main

import (
	"fmt"
	"test-hub-golang/pkg/eventhub"
)

func Consume(consumer eventhub.Consumer) {
	for event := range consumer {
		fmt.Println(event)
	}
}

func main() {
	eventHub := eventhub.New()

	for i := 0; i < 2; i++ {
		consumer := eventHub.Subscribe()
		go Consume(consumer)
	}

	for i := 1; i <= 5; i++ {
		event := fmt.Sprintf("Evento %d", i)
		eventHub.Publish(event)
	}

	eventHub.Close()
}
