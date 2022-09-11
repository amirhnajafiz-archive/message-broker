package main

import "github.com/highway-to-victory/udemy-broker/internal/client"

func main() {
	c, err := client.NewClient("localhost:4040")
	if err != nil {
		panic(err)
	}

	c.Start()
}
