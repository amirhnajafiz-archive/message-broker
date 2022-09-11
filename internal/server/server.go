package server

import (
	"fmt"
	"net"

	"github.com/highway-to-victory/udemy-broker/internal/server/broker"
	"github.com/highway-to-victory/udemy-broker/internal/server/handler"
	"github.com/highway-to-victory/udemy-broker/pkg/logger"
)

func Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	loggerInstance, _ := logger.NewLogger()

	id := 1
	brokerService := broker.NewBroker(loggerInstance)

	go brokerService.Start()

	for {
		conn, er := listener.Accept()
		if er != nil {
			return fmt.Errorf("failed to accept client: %w", er)
		}

		h := handler.NewHandler(id, conn, brokerService.MainChannel, brokerService.TerminateChannel, loggerInstance)

		id++

		brokerService.AddWorker(&h)
		h.Handle()
	}
}
