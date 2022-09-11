package handler

import (
	"log"
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/network"
)

type Handler struct {
	Network network.Network
}

func NewHandler(conn net.Conn) Handler {
	return Handler{
		Network: network.NewNetwork(conn),
	}
}

func (h *Handler) Handle() {
	err := h.Network.Send([]byte("Hello world"))
	if err != nil {
		log.Println(err)
	}
}
