package server

import (
	"fmt"
	"net"

	"github.com/highway-to-victory/udemy-broker/internal/server/broker"
	"github.com/highway-to-victory/udemy-broker/internal/server/handler"
)

func Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	id := 1
	brokerService := broker.NewBroker()

	go brokerService.Start()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept client: %w", err)
		}

		h := handler.NewHandler(conn, brokerService.MainChannel, brokerService.TerminateChannel)
		h.Id = id

		id++

		brokerService.AddWorker(&h)

		h.Handle()
	}
}
