package eventhub

import (
	"errors"
)

type EventHub struct {
	Table map[string]*Topic
}

func New() *EventHub {
	return &EventHub{
		Table: make(map[string]*Topic),
	}
}

func (eh *EventHub) Subscribe(topicChannel string) (*Topic, error) {
	if topic, ok := eh.Table[topicChannel]; !ok {
		return nil, errors.New("channel doesnt exist")
	} else {
		return topic, nil
	}
}

func (eh *EventHub) Publish(topicChannel string, message string) error {
	if topic, ok := eh.Table[topicChannel]; !ok {
		return errors.New("chanel doesnt exist")
	} else {
		topic.Messages <- message
	}

	return nil
}
