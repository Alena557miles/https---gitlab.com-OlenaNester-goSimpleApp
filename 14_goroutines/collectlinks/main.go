package main

import (
	"collectlinks/collect"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var wgResult sync.WaitGroup
	res := make(chan string)
	var result []string

	fileContent, err := os.ReadFile("google.golang.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var n, k int
	k = 5
	n = len(fileContent) / k
	m := make(map[int][]byte)
	for i := 0; i <= len(fileContent); {
		if i == n*k {
			m[i] = fileContent[i:len(fileContent)]
			break
		}
		m[i] = fileContent[i : i+n]
		i = i + n
	}
	c := collect.New()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*0]))
		time.Sleep(time.Second * 2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*1]))
		time.Sleep(time.Second * 2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*2]))
		time.Sleep(time.Second * 2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*3]))
		time.Sleep(time.Second * 2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*4]))
		time.Sleep(time.Second * 2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*5]))
		time.Sleep(time.Second * 2)
	}()

	wgResult.Add(1)
	go func() {
		defer wgResult.Done()
		for i := range res {
			result = append(result, i)
		}
	}()
	wg.Wait()
	close(res)
	wgResult.Wait()

	data := strings.Join(result, " ")
	collect.CreateFileandWriteData("final", data)
	collect.OpenAndPrintFileData("final")

}
