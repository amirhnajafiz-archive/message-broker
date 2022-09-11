package client

import (
	"fmt"
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/logger"
	"github.com/highway-to-victory/udemy-broker/pkg/network"
	"go.uber.org/zap"
)

type Client struct {
	logger *zap.Logger

	network network.Network

	handler func([]byte)
}

func NewClient(address string, handler func([]byte)) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	logInstance, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	var c Client

	c.network = network.NewNetwork(conn)
	c.handler = handler
	c.logger = logInstance

	return &c, nil
}

func (c *Client) Start() {
	go c.listenForDataToGet()
}

func (c *Client) Send(data []byte) error {
	if err := c.network.Send(data); err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

func (c *Client) listenForDataToGet() {
	var buffer = make([]byte, 2048)

	for {
		data, err := c.network.Get(buffer)
		if err != nil {
			c.logger.Error("failed to read", zap.Error(err))
		}

		c.handler(data)
	}
}
