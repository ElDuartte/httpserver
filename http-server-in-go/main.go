package main

import(
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

const keyServerAddr = "serveraddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s\n",
		ctx.Value(keyServerAddr),
		hasFirst, first,
		hasSecond, second)

	io.WriteString(w, "This is my website\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	fmt.Printf("%s: Got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello\n")
}

func startServer(addr string, mux *http.ServeMux, baseCtx context.Context) error {
	server := &http.Server{
		Addr: addr,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.WithValue(baseCtx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed){
		fmt.Printf("sever closed on %s\n", addr)
		return nil
	} else if err != nil {
		return fmt.Errorf("error listening for server %s: %w", addr, err)
	}
		return nil
}

func main(){
	mux := http.NewServeMux() // mux === multiplexer
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx := context.Background()

	go func(){
		if err := startServer("127.0.0.1:6969", mux, ctx); err != nil {
			fmt.Printf("server 1 error: %v\n", err)
		}
	}()

	go func(){
		if err := startServer("127.0.0.1:7070", mux, ctx); err != nil {
			fmt.Printf("server 2 error: %v\n", err)
		}
	}()

	<-ctx.Done()
}
