package eventhub

import "sync"

type Consumer = chan string

type EventHub struct {
	Mutex     sync.Mutex
	Consumers []Consumer
}

func New() *EventHub {
	return &EventHub{
		Consumers: make([]Consumer, 0),
	}
}

func (eh *EventHub) Subscribe() Consumer {
	eh.Mutex.Lock()
	defer eh.Mutex.Unlock()

	consumer := make(Consumer)
	eh.Consumers = append(eh.Consumers, consumer)
	return consumer
}

func (eh *EventHub) Publish(event string) {
	eh.Mutex.Lock()
	defer eh.Mutex.Unlock()

	for _, consumer := range eh.Consumers {
		consumer <- event
	}
}

func (eh *EventHub) Close() {
	eh.Mutex.Lock()
	defer eh.Mutex.Unlock()

	for _, consumer := range eh.Consumers {
		close(consumer)
	}
}
