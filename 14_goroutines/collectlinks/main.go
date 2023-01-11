package main

import (
	"collectlinks/collect"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var result []string

	fileContent, err := os.ReadFile("google.golang.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := collect.New()
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

	var wg sync.WaitGroup
	wg.Add(len(m))
	for _, p := range m {
		go func(p []byte) {
			x := c.FindLinks(string(p))
			result = append(result, x)
			wg.Done()
		}(p)
	}
	wg.Wait()

	data := strings.Join(result, " ")
	collect.CreateFileandWriteData("final", data)
	collect.OpenAndPrintFileData("final")

}
