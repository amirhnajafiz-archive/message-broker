package broker

import "github.com/highway-to-victory/udemy-broker/internal/server/handler"

type Broker struct {
	MainChannel chan []byte
	Handlers    []*handler.Handler
}

func NewBroker() Broker {
	return Broker{
		MainChannel: make(chan []byte),
	}
}

func (b *Broker) Start() {
	for {
		data := <-b.MainChannel

		b.SendData(data)
	}
}

func (b *Broker) AddWorker(h *handler.Handler) {
	b.Handlers = append(b.Handlers, h)
}

func (b *Broker) SendData(data []byte) {
	for _, worker := range b.Handlers {
		if worker.Id != -1 {
			worker.SendChannel <- data
		}
	}
}
