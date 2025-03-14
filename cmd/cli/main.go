package main

import (
	"fmt"
	"sync"
	"test-hub-golang/pkg/eventhub"
	"time"
)

func Consumer(wg *sync.WaitGroup, id int, consumer eventhub.Consumer) {
	defer wg.Done()

	for event := range consumer {
		fmt.Printf("Consumidor %d recebeu a mensagem: %s\n", id, event)
	}
}

func main() {
	var (
		wg       = sync.WaitGroup{}
		eventHub = eventhub.New()
	)

	for i := 0; i < 2; i++ {
		consumer := eventHub.Subscribe()
		wg.Add(1)
		go Consumer(&wg, i, consumer)
	}

	for i := 1; i <= 5; i++ {
		event := fmt.Sprintf("Evento %d", i)
		eventHub.Publish(event)
	}

	go func() {
		<-time.After(time.Second * 2)
		eventHub.Close()
	}()

	wg.Wait()
}
