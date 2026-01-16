package main

import (
	"log"
	"net/http"
)

func main() {
	tracing := true

	var client *http.Client // outer client

	if tracing {
		c, err := createClientWithTracing() //  shadowing happens here
		if err != nil {
			log.Fatal(err)
		}
		client = c
		log.Println("inside if client:", client)
	} else {
		c, err := createDefaultClient() //  shadowing happens here
		if err != nil {
			log.Fatal(err)
		}
		client = c
		log.Println("inside else client:", client)
	}

	// This prints nil because outer client never got assigned
	log.Println("outside client:", client)
}

func createClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}

func createDefaultClient() (*http.Client, error) {
	return &http.Client{}, nil
}
