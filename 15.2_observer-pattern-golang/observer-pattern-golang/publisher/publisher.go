package publisher

import (
	"observer-pattern-golang/subscriber"
)

type Publisher struct {
	subscribers map[string]*subscriber.Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{make(map[string]*subscriber.Subscriber)}
}
func (p *Publisher) AddSubscriber(subscriber *subscriber.Subscriber) {
	p.subscribers[subscriber.Id()] = subscriber
}
func (p *Publisher) RemoveSubscriber(subId string) {
	delete(p.subscribers, subId)
}
func (p *Publisher) Notify(sq float32) {
	for _, subscriber := range p.subscribers {
		subscriber.React(sq)
	}
}

type Subscriber interface {
	React(msg string)
	Id() string
}
