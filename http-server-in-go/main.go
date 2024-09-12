package main

import(
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serveraddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: Got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is root\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	fmt.Printf("%s: Got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello\n")
}

func main(){
	mux := http.NewServeMux() // mux === multiplexer
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx, cancelCtx := context.WithCancel(context.Background())
	
	serverOne := &http.Server{
		Addr: "127.0.0.1:6969",
		Handler: mux,
		BaseContext: func (l net.Listener) context.Context  {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())	
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr: "127.0.0.1:7070",
		Handler: mux,
		BaseContext: func (l net.Listener) context.Context  {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())	
			return ctx
		},
	}

	go func(){
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed){
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func(){
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed){
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
}
