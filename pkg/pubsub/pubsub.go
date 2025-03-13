package pubsub

import (
	"errors"
)

type PubSub struct {
	Table map[string]*Topic
}

func New() *PubSub {
	return &PubSub{
		Table: make(map[string]*Topic),
	}
}

func (ps *PubSub) Subscribe(topicChannel string) (*Topic, error) {
	if topic, ok := ps.Table[topicChannel]; !ok {
		return nil, errors.New("channel doesnt exist")
	} else {
		return topic, nil
	}
}

func (ps *PubSub) Publish(topicChannel string, message string) error {
	if topic, ok := ps.Table[topicChannel]; !ok {
		return errors.New("chanel doesnt exist")
	} else {
		topic.Messages <- message
	}

	return nil
}
