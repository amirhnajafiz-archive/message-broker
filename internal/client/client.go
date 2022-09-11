package client

import (
	"fmt"
	"net"

	"github.com/highway-to-victory/udemy-broker/pkg/logger"
	"github.com/highway-to-victory/udemy-broker/pkg/network"
	"go.uber.org/zap"
)

// Client manages the data transferring in client side.
type Client struct {
	// logger instance.
	logger *zap.Logger
	// enabled is for executing handler function.
	enabled bool
	// network manages the requests over tcp.
	network network.Network
	// handler is the handling function.
	handler func([]byte)
}

// NewClient builds a new client and connects to broker server.
func NewClient(address string, handler func([]byte)) (*Client, error) {
	// connect to server
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	// create logger instance
	logInstance, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	// create client instance
	var c Client

	c.network = network.NewNetwork(conn)
	c.handler = handler
	c.logger = logInstance
	c.enabled = true

	return &c, nil
}

// Start will start the client.
func (c *Client) Start() {
	go c.listenForDataToGet()
}

// Send a data to server.
func (c *Client) Send(data []byte) error {
	if err := c.network.Send(data); err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

// Enable the handler.
func (c *Client) Enable() {
	c.enabled = true
}

// Disable the handler.
func (c *Client) Disable() {
	c.enabled = false
}

// this method waits for incoming data
// and executes the handler function.
func (c *Client) listenForDataToGet() {
	var buffer = make([]byte, 2048)

	for {
		data, err := c.network.Get(buffer)
		if err != nil {
			c.logger.Error("failed to read", zap.Error(err))
		}

		if c.enabled {
			c.handler(data)
		}
	}
}
