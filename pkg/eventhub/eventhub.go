package eventhub

type Consumer = chan string

type EventHub struct {
	Consumers []Consumer
}

func New() *EventHub {
	return &EventHub{
		Consumers: make([]Consumer, 0),
	}
}

func (eh *EventHub) Subscribe() Consumer {
	consumer := make(Consumer)
	eh.Consumers = append(eh.Consumers, consumer)
	return consumer
}

func (eh *EventHub) Publish(event string) {
	for _, consumer := range eh.Consumers {
		consumer <- event
	}
}

func (eh *EventHub) Close() {
	for _, consumer := range eh.Consumers {
		close(consumer)
	}
}
