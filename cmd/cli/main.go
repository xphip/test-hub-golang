package main

import (
	"fmt"
	"test-hub-golang/pkg/eventhub"
)

func Consumer(id int, consumer eventhub.Consumer) {
	for event := range consumer {
		fmt.Printf("Consumidor %d recebeu a mensagem: %s\n", id, event)
	}
}

func main() {
	eventHub := eventhub.New()

	for i := 0; i < 2; i++ {
		consumer := eventHub.Subscribe()
		go Consumer(i, consumer)
	}

	for i := 1; i <= 5; i++ {
		event := fmt.Sprintf("Evento %d", i)
		eventHub.Publish(event)
	}

	eventHub.Close()
}
