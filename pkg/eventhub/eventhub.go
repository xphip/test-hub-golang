package eventhub

import "sync"

type Consumer = chan string

type EventHub struct {
	Mutex     sync.Mutex
	Closed    bool
	Consumers []Consumer
}

func New() *EventHub {
	return &EventHub{
		Closed:    false,
		Consumers: make([]Consumer, 0),
	}
}

func (eh *EventHub) Subscribe() Consumer {
	eh.Mutex.Lock()
	defer eh.Mutex.Unlock()

	if eh.Closed {
		return nil
	}

	consumer := make(Consumer)
	eh.Consumers = append(eh.Consumers, consumer)
	return consumer
}

func (eh *EventHub) Publish(event string) {
	eh.Mutex.Lock()
	defer eh.Mutex.Unlock()

	if eh.Closed {
		return
	}

	for _, consumer := range eh.Consumers {
		go func(c Consumer) {
			c <- event
		}(consumer)
	}
}

func (eh *EventHub) Close() {
	eh.Mutex.Lock()
	defer eh.Mutex.Unlock()

	if eh.Closed {
		return
	}

	eh.Closed = true

	for _, consumer := range eh.Consumers {
		close(consumer)
	}
}
