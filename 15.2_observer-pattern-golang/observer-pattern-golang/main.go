package main

import (
	"observer-pattern-golang/subject"
	"observer-pattern-golang/subscriber"
)

func main() {
	a := subject.OccupiedTerritory("Ukraine", 120.7)

	s1 := subscriber.NewSubscriber("Ivan")
	s2 := subscriber.NewSubscriber("Svetlana")
	a.AddSubscriber(s1)
	a.AddSubscriber(s2)

	a.FreeTerritory(120)
	//a.FreeTerritory(0.7)

	s3 := subscriber.NewSubscriber("Oleksiy")
	a.AddSubscriber(s3)
	a.RemoveSubscriber("Ivan")
	a.FreeTerritory(200)

}
