package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test-hub-golang/pkg/eventhub"
)

type CLI struct {
	Ctx      context.Context
	EventHub *eventhub.EventHub
}

func NewCli() *CLI {
	return &CLI{EventHub: eventhub.New()}
}

func (cli *CLI) Run() error {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("EventHub started")

	// TODO: make listener service to handle goroutines

	<-wait
	fmt.Println("EventHub finished")
	return nil
}
