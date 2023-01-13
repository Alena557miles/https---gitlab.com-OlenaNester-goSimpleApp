package subscriber

import "log"

type Subscriber struct {
	subId string
}

func NewSubscriber(subId string) *Subscriber {
	return &Subscriber{subId: subId}
}
func (s *Subscriber) React(sq float32) {
	log.Printf("Subscriber with id '%s' get information about territory's changes: %4.2f", s.subId, sq)
}
func (s *Subscriber) Id() string {
	return s.subId
}
