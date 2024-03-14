package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type requestJson struct {
	Primero int `json:"primero,omitempty"`
	Segundo int `json:"segundo,omitempty"`
}

type responseJson struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

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
		result := fmt.Sprintf("La suma es: %d", numConverted+numConverted2)
		writer.Write([]byte(result))
		writer.WriteHeader(200)

	})

	http.HandleFunc("/resta", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		resta := requestJson{}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()

		err = json.Unmarshal(body, &resta)
		if err != nil {
			resJson, err := json.Marshal(responseJson{
				Message: fmt.Sprintf("Hubo un error en Unmarshal, el error es %v", err),
				Status:  http.StatusInternalServerError,
			})
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			writer.Write(resJson)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		result := resta.Primero - resta.Segundo
		send := fmt.Sprintf("La resta es: %d", result)

		res := responseJson{
			Message: send,
			Status:  http.StatusOK,
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(resJson)
		writer.WriteHeader(http.StatusAccepted)
	})

	log.Fatal(http.ListenAndServe(":3030", nil))
}
