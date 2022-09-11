package broker

import (
	"github.com/highway-to-victory/udemy-broker/internal/server/handler"
	"go.uber.org/zap"
)

// Broker
// is the main broker service.
type Broker struct {
	// logger instance
	logger *zap.Logger
	// communication channels
	MainChannel      chan []byte
	TerminateChannel chan int
	// list of the Handlers
	Handlers []*handler.Handler
}

// NewBroker generates a new broker.
func NewBroker(logger *zap.Logger) Broker {
	return Broker{
		logger:           logger,
		MainChannel:      make(chan []byte),
		TerminateChannel: make(chan int),
	}
}

// Start broker service.
func (b *Broker) Start() {
	for {
		select {
		case data := <-b.MainChannel:
			// sending data
			b.sendData(data)
		case id := <-b.TerminateChannel:
			// removing worker
			b.removeWorker(id)
		}
	}
}

// AddWorker to list.
func (b *Broker) AddWorker(h *handler.Handler) {
	b.Handlers = append(b.Handlers, h)
}

func (b *Broker) removeWorker(id int) {
	for index, worker := range b.Handlers {
		if worker.Id == id {
			b.Handlers = append(b.Handlers[:index], b.Handlers[index+1:]...)

			b.logger.Info("worker terminated", zap.Int("id", id))

			break
		}
	}
}

func (b *Broker) sendData(data []byte) {
	for _, worker := range b.Handlers {
		worker.SendChannel <- data
	}
}
