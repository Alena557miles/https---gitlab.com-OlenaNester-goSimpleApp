package publisher

import "observer-pattern-golang/subscriber"

type Publisher struct {
	subscribers map[string]*subscriber.Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{make(map[string]*subscriber.Subscriber)}
}
func (p *Publisher) AddSubscribers(subscriber *subscriber.Subscriber) {
	p.subscribers[subscriber.Id()] = subscriber
}
func (p *Publisher) RemoveSubscriber(subId string) {
	delete(p.subscribers, subId)
}
func (p *Publisher) Notify(msg string) {
	for _, subscriber := range p.subscribers {
		subscriber.React(msg)
	}
}

type Subscriber interface {
	React(msg string)
	Id() string
}
