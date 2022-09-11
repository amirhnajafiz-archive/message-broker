package broker

import (
	"github.com/highway-to-victory/udemy-broker/internal/server/handler"
)

type Broker struct {
	MainChannel      chan []byte
	TerminateChannel chan int
	Handlers         []*handler.Handler
}

func NewBroker() Broker {
	return Broker{
		MainChannel:      make(chan []byte),
		TerminateChannel: make(chan int),
	}
}

func (b *Broker) Start() {
	for {
		select {
		case data := <-b.MainChannel:
			b.SendData(data)
		case id := <-b.TerminateChannel:
			b.RemoveWorker(id)
		}
	}
}

func (b *Broker) AddWorker(h *handler.Handler) {
	b.Handlers = append(b.Handlers, h)
}

func (b *Broker) RemoveWorker(id int) {
	for index, worker := range b.Handlers {
		if worker.Id == id {
			b.Handlers = append(b.Handlers[:index], b.Handlers[index+1:]...)

			break
		}
	}
}

func (b *Broker) SendData(data []byte) {
	for _, worker := range b.Handlers {
		if worker.Id != -1 {
			worker.SendChannel <- data
		}
	}
}
