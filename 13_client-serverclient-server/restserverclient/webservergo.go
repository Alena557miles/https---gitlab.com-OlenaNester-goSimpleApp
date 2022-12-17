package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func doSomething(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var number string = vars["number"]
	resp := make(map[string]string)
	resp["message"] = `recieved number is :  ` + number
	number = number + "1"
	log.Printf(`Value after concat is : %s`, number)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
	startServer(number)
}

func startServer(num any) {
	//log.Printf(`number from first server: %s`, num)
	c := http.Client{Timeout: time.Second}
	req, err := http.NewRequest(`POST`, `http://localhost:8081/concat`, nil)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	body := []byte(`{	
		"jsonrpc": 2.0,
		"method": "concat.ConcatNumbers",
		"params": [{
		"Content": "7"
		}]
	})`)
	req.Header.Add(`Content-Type`, `application/json`)
	neededBody := io.NopCloser(bytes.NewReader(body))
	req.Body = neededBody
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	ctx := context.Background()
	router := mux.NewRouter()
	srv := &http.Server{
		Addr:              `0.0.0.0:5000`,
		ReadTimeout:       time.Millisecond * 200,
		WriteTimeout:      time.Millisecond * 200,
		IdleTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Millisecond * 200,
		Handler:           router,
	}
	// "localhost:5000/{number}"
	router.HandleFunc("/{number}", doSomething)

	go func() {
		log.Println(`Web Server started`)
		err := srv.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	<-done

	log.Println(`Web Server is shutting down`)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(ctx, err)
	}

}