package main

import "github.com/highway-to-victory/udemy-broker/internal/server"

// testing our broker service.
func main() {
	// creating a broker server on 4040
	if err := server.Start(":4040"); err != nil {
		panic(err)
	}
}
