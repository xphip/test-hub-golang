package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test-hub-golang/pkg/eventhub"
)

func Consume(consumer eventhub.Consumer) {
	for event := range consumer {
		fmt.Println(event)
	}
}

func main() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

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

	<-wait
	fmt.Println("EventHub finished")
}
