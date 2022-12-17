package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

type SomeService struct{}
type SomeArgs struct {
	Receiver, Content string
}
type Response struct {
	Result string
}

func (s *SomeService) ConcatNumbers(r *http.Request, args *SomeArgs, result *Response) error {
	log.Printf(`Value from first server: %s`, args.Content)
	var res string
	res = args.Content + "2"
	*result = Response{Result: res}
	log.Printf(`Value after concat is: %s `, res)
	return nil
}

func main() {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), `application/json`)

	concat := new(SomeService)

	_ = rpcServer.RegisterService(concat, `concat`)

	router := mux.NewRouter()
	router.Handle(`/concat`, rpcServer)
	log.Fatal(http.ListenAndServe(`:8081`, router))
}
