package main

import (
	"context"

	"github.com/osniantonio/goexpert/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// rabbitmq.Publish(ch, "Hello World!", "amq.direct")

	ctx := context.Background()

	rabbitmq.PublishWithContext(ctx, ch, "Hello World!", "amq.direct")

}
