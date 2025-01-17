package main

import (
	"context"

	"github.com/rl404/fairy/pubsub"
)

type sampleData struct {
	Field1 string
	Field2 int
}

func main() {
	// Pubsub type.
	t := pubsub.RabbitMQ

	// Init client.
	client, err := pubsub.New(t, "amqp://guest:guest@localhost:5672", "")
	if err != nil {
		panic(err)
	}

	// Don't forget to close.
	defer client.Close()

	// Sample data. Can be any type.
	data := sampleData{
		Field1: "a",
		Field2: 1,
	}

	// Publish data to specific topic/channel. Data will be encoded first.
	if err = client.Publish(context.Background(), "topic", data); err != nil {
		panic(err)
	}
}
