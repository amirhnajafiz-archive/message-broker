package main

import "github.com/highway-to-victory/udemy-broker/internal/server"

func main() {
	if err := server.Start(":4040"); err != nil {
		panic(err)
	}
}
