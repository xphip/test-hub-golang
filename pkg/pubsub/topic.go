package pubsub

type Topic struct {
	Messages  chan string
	Consumers []chan string
}

func (topic *Topic) Publish(message string) {
	for _, consumer := range topic.Consumers {
		consumer <- message
	}
}
