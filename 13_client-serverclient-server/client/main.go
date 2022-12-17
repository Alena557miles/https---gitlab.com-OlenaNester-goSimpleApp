package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	req, err := http.NewRequest(`POST`, `http://localhost:5000/6`, nil)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	defer resp.Body.Close()
}
