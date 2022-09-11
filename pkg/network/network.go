package network

import (
	"fmt"
	"io"
	"net"
)

type Network struct {
	connection net.Conn
}

func NewNetwork(conn net.Conn) Network {
	return Network{
		connection: conn,
	}
}

func (n *Network) Send(data []byte) error {
	if _, err := n.connection.Write(data); err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

func (n *Network) Get(buffer []byte) ([]byte, error) {
	bytes, err := n.connection.Read(buffer)
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf("failed to get data: %w", err)
		}
	}

	return buffer[:bytes], nil
}
