package main

import(
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / request\n")
	io.WriteString(w, "This is my web\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, http\n")
}

func main(){
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe("127.0.0.1:6969", nil)
	if errors.Is(err, http.ErrServerClosed){
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
