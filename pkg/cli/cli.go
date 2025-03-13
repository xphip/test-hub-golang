package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test-hub-golang/pkg/pubsub"
)

type CLI struct {
	Ctx    context.Context
	PubSub *pubsub.PubSub
}

func NewCli() *CLI {
	return &CLI{PubSub: pubsub.New()}
}

func (cli *CLI) Run() error {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("PubSub started")

	// TODO: make listener service to handle goroutines

	<-wait
	fmt.Println("PubSub finished")
	return nil
}
