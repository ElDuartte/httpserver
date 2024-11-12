package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serveraddr"

func serveHTML(w http.ResponseWriter, r *http.Request, filePath string){
	if _, err := os.Stat(filePath); os.IsNotExist(err){
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filePath)
}

func getRoot(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	cookie, err := r.Cookie("my-cookie")
	if err == nil{
		fmt.Printf("Cookie value: %s\n", cookie.Value)
	}

	setCookie(w, "my-cookie", "cookie-value-test")

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))

	serveHTML(w, r, "./home.html")
}

func getHello(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	fmt.Printf("%s: Got /hello request\n", ctx.Value(keyServerAddr))

	serveHTML(w, r, "./show.html")
}

func startServer(addr string, handler http.Handler, baseCtx context.Context) error {
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(baseCtx, keyServerAddr, l.Addr().String())
		},
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("error listening for server %s: %w", addr, err)
	}
	fmt.Printf("server closed on %s\n", addr)
	return nil
}

func setCookie(w http.ResponseWriter, name, value string){
	cookie := &http.Cookie{
		Name: name,
		Value: value,
		Path: "/",
	}
	http.SetCookie(w, cookie)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/show", getHello)

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

