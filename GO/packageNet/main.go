package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HOLAAAAA")

	welcomeHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Bienvenido!\n")
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")

	}

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
