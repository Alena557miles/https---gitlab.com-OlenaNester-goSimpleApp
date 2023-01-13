package main

import (
	"observer-pattern-golang/publisher"
	"observer-pattern-golang/subscriber"
)

func main() {
	p := publisher.NewPublisher()

	s := subscriber.NewSubscriber("23")
	p.AddSubscribers(s)
	p.Notify("sun is shining")

	s1 := subscriber.NewSubscriber("4")
	p.AddSubscribers(s1)
	p.Notify("hello")

	s3 := subscriber.NewSubscriber("67")
	p.AddSubscribers(s3)
	p.RemoveSubscriber("23")

	p.Notify("bye")

}
