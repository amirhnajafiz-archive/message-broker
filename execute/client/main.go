package main

import (
	"os"

	"github.com/highway-to-victory/udemy-broker/internal/client"
)

func main() {
	c, err := client.NewClient("localhost:4040")
	if err != nil {
		panic(err)
	}

	if os.Args[1] == "pub" {
		c.Send([]byte("Hello from client"))
	} else {
		c.Start()

		select {}
	}
}
