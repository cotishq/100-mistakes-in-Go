package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	tracing := true

	fmt.Println("========== BUGGY VERSION ==========")
	buggy(tracing)

	fmt.Println("\n========== FIXED VERSION ==========")
	fixed(tracing)
}

//  Bug: unintended variable shadowing using :=
func buggy(tracing bool) {
	var client *http.Client

	if tracing {
		client, err := createClientWithTracing() //  shadows outer client
		if err != nil {
			log.Println("error:", err)
			return
		}
		log.Println("inside if:", client)
	} else {
		client, err := createDefaultClient() //  shadows outer client
		if err != nil {
			log.Println("error:", err)
			return
		}
		log.Println("inside else:", client)
	}

	//  Outer client never got assigned -> nil
	log.Println("outside:", client)
}

//  Fix: declare err once + use "=" assignment (no shadowing)
func fixed(tracing bool) {
	var client *http.Client
	var err error

	if tracing {
		client, err = createClientWithTracing() //  no :=
	} else {
		client, err = createDefaultClient()
	}

	if err != nil {
		log.Println("common error handling:", err)
		return
	}

	//  Outer client is assigned properly
	log.Println("outside:", client)
}

func createClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}

func createDefaultClient() (*http.Client, error) {
	return &http.Client{}, nil
}
