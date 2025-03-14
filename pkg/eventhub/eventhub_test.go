package eventhub_test

import (
	"sync"
	"test-hub-golang/pkg/eventhub"
	"testing"
)

func TestSubscribeAndPublish(t *testing.T) {
	var (
		wg                = sync.WaitGroup{}
		eventHub          = eventhub.New()
		consumer          = eventHub.Subscribe()
		eventTest         = "This is and event test"
		eventTestReceived string
	)

	wg.Add(1)

	go func(g *sync.WaitGroup, c eventhub.Consumer) {
		defer g.Done()
		eventTestReceived = <-c
	}(&wg, consumer)

	eventHub.Publish(eventTest)

	wg.Wait()

	if eventTest != eventTestReceived {
		t.Error("event received doesnt match with espected event")
	}
}
