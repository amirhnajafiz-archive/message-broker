package server

import (
	"fmt"
	"net"

	"github.com/highway-to-victory/udemy-broker/internal/server/handler"
)

func Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept client: %w", err)
		}

		fmt.Println("accept client")

		h := handler.NewHandler(conn)

		go h.Handle()
	}
}
