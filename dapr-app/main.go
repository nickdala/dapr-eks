package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"log"
	"net/http"
	"strconv"
)

const (
	stateStoreName = "statestore"
	daprPort       = "3500"
)

var daprClient dapr.Client

func init() {
	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	daprClient = client
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	fmt.Printf("Getting greetings\n")
	item, err := daprClient.GetState(ctx, stateStoreName, "greetings")

	if err != nil {
		fmt.Printf("Failed to get state: %v\n", err)
	}

	var numberOfGreetings = 0
	if len(item.Value) > 0 {
		fmt.Printf("Value of greetings from redis %s\n", item.Value)
		numberOfGreetings, _ = strconv.Atoi(string(item.Value))
	} else {
		fmt.Printf("Greetings not found\n")
	}
	numberOfGreetings++

	fmt.Printf("Saving greetings %d\n", numberOfGreetings)
	if err := daprClient.SaveState(ctx, stateStoreName, "greetings", []byte(strconv.Itoa(numberOfGreetings))); err != nil {
		fmt.Printf("Failed to persist state: %v\n", err)
	} else {
		fmt.Printf("Successfully persisted state\n")
	}

	fmt.Fprintf(w, "Greetings called %d times!\n", numberOfGreetings)
}

func main() {
	// properly close the dapr client
	defer daprClient.Close()

	http.HandleFunc("/api/v1/message", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
