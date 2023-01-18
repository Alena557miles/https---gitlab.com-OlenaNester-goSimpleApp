package subject

import (
	"log"
	"observer-pattern-golang/publisher"
)

type Area struct {
	name string
	sq   float32
	*publisher.Publisher
}

func OccupiedTerritory(name string, sq float32) *Area {
	log.Printf("Country: %s, occupied territory: %4.2f", name, sq)
	sub := publisher.NewPublisher()
	return &Area{
		name,
		sq,
		sub,
	}
}

func (a *Area) FreeTerritory(sq float32) {
	if (a.sq >= 0) && (a.sq-sq > 0) {
		a.sq = a.sq - sq
		log.Printf("For now liberated territory is: %4.2f", a.sq)
		a.Notify(a.sq)
	} else if (a.sq > 0) && (a.sq-sq < 0) {
		a.sq = sq - a.sq
		log.Printf("For now country %s are expanded its territory for: %4.2f", a.name, a.sq)
		a.Notify(a.sq)
	} else {
		a.sq = a.sq + sq
		log.Printf("For now country %s are expanded its territory for: %4.2f", a.name, a.sq)
		a.Notify(a.sq)
	}
}

type Subscriber interface {
	React(msg float32)
	Id() string
}

type Publisher interface {
	AddSubscriber(Subscriber)
	RemoveSubscriber(subId string)
}
