package server

import (
	"fmt"
	"net"

	"github.com/highway-to-victory/udemy-broker/internal/server/broker"
	"github.com/highway-to-victory/udemy-broker/internal/server/handler"
	"github.com/highway-to-victory/udemy-broker/pkg/logger"
)

// Start our broker server.
func Start(address string) error {
	// creating a listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	// creat a logger instance
	loggerInstance, _ := logger.NewLogger()

	// initializing a broker service
	id := 1
	brokerService := broker.NewBroker(loggerInstance)

	// starting broker serivce
	go brokerService.Start()

	for {
		// accept clients
		conn, er := listener.Accept()
		if er != nil {
			return fmt.Errorf("failed to accept client: %w", er)
		}

		// create a new handler (worker)
		h := handler.NewHandler(id, conn, brokerService.MainChannel, brokerService.TerminateChannel, loggerInstance)

		id++

		// execute worker
		brokerService.AddWorker(&h)
		h.Handle()
	}
}
