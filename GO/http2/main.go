package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//Creación del http Handler por defecto, también hay por custom
	http.HandleFunc("/saludo", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			//writer.Write([]byte("Metodo no permitido"))
			writer.WriteHeader(405)
			return
		}

		nombre := request.URL.Query().Get("nombre")
		edad := request.URL.Query().Get("edad")

		saludo := fmt.Sprintf("Hola %s tu edad es %s", nombre, edad)
		writer.Write([]byte(saludo))
		writer.WriteHeader(200)
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		saludo := fmt.Sprintf("inicio")
		writer.Write([]byte(saludo))
	})

	http.HandleFunc("/suma", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			writer.WriteHeader(405)
			return
		}

		num1 := request.URL.Query().Get("primero")
		num2 := request.URL.Query().Get("segundo")
		//Strconv = string converter --> PAQUETE en GO para convertir tipos de datos a otro
		numConverted, err := strconv.Atoi(num1)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		numConverted2, err := strconv.Atoi(num2)
		if err != nil {
			//Sustitución del código en num entero por el status del código directamente
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		//Unión/Concatenación de la suma de los dos números ya convertidos de tipo string a int
		saludo := fmt.Sprintf("La suma es: %d", numConverted+numConverted2)
		writer.Write([]byte(saludo))
		writer.WriteHeader(200)

	})

	log.Fatal(http.ListenAndServe(":3030", nil))
}
