package main

import (
	"log"

	"github.com/kelseyhightower/google-cloud-functions-go/event"
)

func F(e event.Event) (string, error) {
	log.Printf("processing event: %s", e.EventID)
	return "", nil
}
