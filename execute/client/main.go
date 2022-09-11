package main

import (
	"log"
	"os"

	"github.com/highway-to-victory/udemy-broker/internal/client"
)

// testing our clients.
func main() {
	// creating a new client and connect to server on port 4040
	c, err := client.NewClient("localhost:4040", func(bytes []byte) {
		// handler function
		log.Printf("client got: %s\n", string(bytes))
	})
	if err != nil {
		panic(err)
	}

	// check if the client should publish or not
	if os.Args[1] == "pub" {
		c.Disable()
		_ = c.Send([]byte("Hello from client"))
	} else {
		c.Start()
		select {}
	}
}
