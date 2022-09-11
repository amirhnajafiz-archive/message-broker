package main

import (
	"log"
	"os"

	"github.com/highway-to-victory/udemy-broker/internal/client"
)

func main() {
	c, err := client.NewClient("localhost:4040", func(bytes []byte) {
		log.Printf("client got: %s\n", string(bytes))
	})
	if err != nil {
		panic(err)
	}

	if os.Args[1] == "pub" {
		_ = c.Send([]byte("Hello from client"))
	} else {
		c.Start()

		select {}
	}
}
