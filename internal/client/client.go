package client

import (
	"fmt"
	"log"
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/network"
)

type Client struct {
	Network network.Network
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	var c Client

	c.Network = network.NewNetwork(conn)

	return &c, nil
}

func (c *Client) Start() {
	go c.listenForGetData()
}

func (c *Client) Send(data []byte) error {
	if err := c.Network.Send(data); err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

func (c *Client) listenForGetData() {
	var buffer = make([]byte, 2048)

	for {
		data, err := c.Network.Get(buffer)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(data))
	}
}
