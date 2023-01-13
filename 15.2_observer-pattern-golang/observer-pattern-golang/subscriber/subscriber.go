package subscriber

import "log"

type Subscriber struct {
	subId string
}

func NewSubscriber(subId string) *Subscriber {
	return &Subscriber{subId: subId}
}
func (s *Subscriber) React(msg string) {
	log.Printf("Subscriber with id %s get message: %s", s.subId, msg)
}
func (s *Subscriber) Id() string {
	return s.subId
}
