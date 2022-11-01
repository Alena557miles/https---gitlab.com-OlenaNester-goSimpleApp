package main

import "fmt"

type ServiceLocator struct {
	services []any
	names    []string
}

func (s *ServiceLocator) Register(name string, some any) {
	s.services = append(s.services, some)
	s.names = append(s.names, name)
}

func (s *ServiceLocator) Get(name string) bool {
	for _, t := range s.names {
		if t == name {
			return true
		}
	}
	return false
}

type Concrete struct {
	Name string
}

func (c *Concrete) Do() {}

type Doer interface {
	Do()
}

func main() {
	l := ServiceLocator{}
	l.Register(`serv1`, &Concrete{"tomate"})

	if l.Get(`serv1`) {
		fmt.Println("by interface pointer ok")
	}

	// var z Doer
	// if l.Get(&z) {
	// 	fmt.Println("by interface ok")
	// 	fmt.Println(z)
	// }

	// var y *Concrete
	// if l.Get(&y) {
	// 	fmt.Println("by struct pointer ok")
	// 	fmt.Println(y)
	// }
}
