package handler

import (
	"log"
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/network"
)

type Handler struct {
	Id               int
	GetChannel       chan []byte
	SendChannel      chan []byte
	TerminateChannel chan int
	Network          network.Network
}

func NewHandler(conn net.Conn, channel chan []byte, terminateChannel chan int) Handler {
	return Handler{
		GetChannel:       channel,
		SendChannel:      make(chan []byte),
		TerminateChannel: terminateChannel,
		Network:          network.NewNetwork(conn),
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
			h.terminate()

			break
		}

		h.GetChannel <- data
	}
}

func (h *Handler) terminate() {
	h.TerminateChannel <- h.Id

	log.Printf("worker termiated: %d\n", h.Id)
}
