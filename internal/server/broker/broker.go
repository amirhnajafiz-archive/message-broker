package broker

import (
	"github.com/highway-to-victory/udemy-broker/internal/server/handler"
	"go.uber.org/zap"
)

type Broker struct {
	logger *zap.Logger

	MainChannel      chan []byte
	TerminateChannel chan int

	Handlers []*handler.Handler
}

func NewBroker(logger *zap.Logger) Broker {
	return Broker{
		logger:           logger,
		MainChannel:      make(chan []byte),
		TerminateChannel: make(chan int),
	}
}

func (b *Broker) Start() {
	for {
		select {
		case data := <-b.MainChannel:
			b.sendData(data)
		case id := <-b.TerminateChannel:
			b.removeWorker(id)
		}
	}
}

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
