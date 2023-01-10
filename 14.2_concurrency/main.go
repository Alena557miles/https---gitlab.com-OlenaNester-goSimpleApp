package main

import (
	"fmt"
	"log"
	"sync"
)

type Collect struct {
	_arrayOfString []string
	mu             sync.RWMutex
}

func (c *Collect) makeArray(n int, arr []string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return arr[n]
}

func main() {

	s := []string{"a", "b", "c", "d", "e"}

	c := &Collect{}

	var result []string
	var wg sync.WaitGroup
	var wgRes sync.WaitGroup
	res := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.makeArray(0, s)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.makeArray(1, s)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.makeArray(2, s)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.makeArray(3, s)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.makeArray(4, s)
	}()

	wgRes.Add(1)
	go func() {
		defer wgRes.Done()
		for i := range res {
			fmt.Println(i)
			result = append(result, i)
		}
	}()
	wg.Wait()
	close(res)
	wgRes.Wait()
	log.Println(result)
}
