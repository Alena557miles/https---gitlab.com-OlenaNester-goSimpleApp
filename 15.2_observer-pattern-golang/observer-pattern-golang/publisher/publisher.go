package publisher

type Publisher struct {
	subscribers map[string]Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{make(map[string]Subscriber)}
}
func (p *Publisher) AddSubscriber(subscriber Subscriber) {
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
	React(msg float32)
	Id() string
}
