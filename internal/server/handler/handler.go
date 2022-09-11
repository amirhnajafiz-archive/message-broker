package handler

import (
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/network"
	"go.uber.org/zap"
)

type Handler struct {
	Id int

	logger *zap.Logger

	GetChannel       chan []byte
	SendChannel      chan []byte
	TerminateChannel chan int

	Network network.Network
}

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
		Network:          network.NewNetwork(conn),
		logger:           logger,
	}
}

func (h *Handler) Handle() {
	go h.listenForDataToSend()
	go h.listenForDataToGet()
}

func (h *Handler) listenForDataToSend() {
	for {
		message := <-h.SendChannel

		err := h.Network.Send(message)
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
		data, err := h.Network.Get(buffer)

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
