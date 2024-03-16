package main

import (
	"fmt"

	"github.com/osniantonio/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
