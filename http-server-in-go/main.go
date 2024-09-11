package main

import(
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWritter, r *http.Request) {
	fmt.Println("Got / request\n")
	io.WriteString(w, "This is my web\n")
}

func getHello(w http.ResponseWritter, r *http.Request){
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, http\n")
}
