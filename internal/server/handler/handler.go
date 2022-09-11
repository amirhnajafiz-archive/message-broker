package handler

import (
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/network"
	"go.uber.org/zap"
)

// Handler manages to keep the client connected
// to broker service.
type Handler struct {
	// worker id
	Id int
	// logger instance
	logger *zap.Logger
	// communication channels
	GetChannel       chan []byte
	SendChannel      chan []byte
	TerminateChannel chan int
	// network instance
	network network.Network
}

// NewHandler creates a new worker.
func NewHandler(
	id int,
	conn net.Conn,
	channel chan []byte,
	terminateChannel chan int,
	logger *zap.Logger,
) Handler {
	return Handler{
		Id:               id,
		GetChannel:       channel,
		SendChannel:      make(chan []byte),
		TerminateChannel: terminateChannel,
		network:          network.NewNetwork(conn),
		logger:           logger,
	}
}

// Handle
// starts handling the client.
func (h *Handler) Handle() {
	go h.listenForDataToSend()
	go h.listenForDataToGet()
}

func (h *Handler) listenForDataToSend() {
	for {
		message := <-h.SendChannel

		err := h.network.Send(message)
		if err != nil {
			h.logger.Error("worker error", zap.Error(err))

			h.terminate()

			break
		}
	}
}

func (h *Handler) listenForDataToGet() {
	var buffer = make([]byte, 2048)

	for {
		data, err := h.network.Get(buffer)

		if err != nil {
			h.logger.Error("worker error", zap.Error(err))

			h.terminate()

			break
		}

		h.GetChannel <- data
	}
}

func (h *Handler) terminate() {
	h.TerminateChannel <- h.Id
}
