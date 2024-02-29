package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// https://pkg.go.dev/net/http

func main() {
	/*
		response, err := http.Get("https://www.upemor.edu.mx/")
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	*/

	// servidor

	// /saludo
	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola me llamaste desde: %q con el metodo: %s", html.EscapeString(r.URL.Path), r.Method)
		log.Println("Se recibió un request de:", r.Method, r.URL.User)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
