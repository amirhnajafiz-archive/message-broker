package handler

import (
	"log"
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/network"
)

type Handler struct {
	Id          int
	GetChannel  chan []byte
	SendChannel chan []byte
	Network     network.Network
}

func NewHandler(conn net.Conn, channel chan []byte) Handler {
	return Handler{
		GetChannel:  channel,
		SendChannel: make(chan []byte),
		Network:     network.NewNetwork(conn),
	}
}

func (h *Handler) Handle() {
	go h.listenForDataToSend()
	h.listenForDataToGet()

	h.Id = -1
}

func (h *Handler) listenForDataToSend() {
	for {
		message := <-h.SendChannel

		err := h.Network.Send(message)
		if err != nil {
			log.Println(err)

			break
		}
	}
}

func (h *Handler) listenForDataToGet() {
	var buffer = make([]byte, 2048)

	for {
		data, err := h.Network.Get(buffer)

		if err != nil {
			log.Println(err)

			continue
		}

		h.GetChannel <- data
	}
}
