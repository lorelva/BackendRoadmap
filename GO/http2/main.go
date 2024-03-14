package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Para el endpoint de "/resta",request y response
type requestJson struct {
	Primero int `json:"primero,omitempty"`
	Segundo int `json:"segundo,omitempty"`
}

type responseJson struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// Para el endpoint de "/name", request y repsonse
type requestNameJson struct {
	Nombre   string `json:"nombre,omitempty"`
	Apellido string `json:"apellido,omitempty"`
	Edad     int    `json:"edad,omitempty"`
}

type responseNameJson struct {
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

	//Uso de Marshal y Unmarshal, para formato JSON en endpoint "/resta"
	http.HandleFunc("/resta", func(writer http.ResponseWriter, request *http.Request) {
		//Verificación del método HTTP deseado
		if request.Method != "POST" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		resta := requestJson{}

		//Lectura del body del request
		//body es de tipo io.ReadAll, y al final cerrar el body del request con defer
		body, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()

		//json.Unmarshal()= recibir JSON para convertir en una estructura de datos Go.
		err = json.Unmarshal(body, &resta)
		if err != nil {
			//happy Path--> hacer que falle
			//responder como API
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

		//json.Marshal()= Estructura de datos Go para convertir en JSON.
		resJson, err := json.Marshal(res)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(resJson)
		//Metódo Write del http.ResponseWriter , escribe en el body del response(JSON)-->
		//Para enviar un mensaje de respuesta al cliente en el cuerpo de la respuesta HTTP
		//Ya convertido(casteo) el string(mensaje) a []bytes = cadena/arreglo de bytes
		//writer.Write([]byte(send))
		writer.WriteHeader(http.StatusAccepted)
	})

	http.HandleFunc("/nombre", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		//Recibir JSON = Unmarshal
		saludo := requestNameJson{}

		//Leer el body del request
		body, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		//Cerrar el body del request
		defer request.Body.Close()

		//1er parámetro = bytes que contiene datos en formato JSON
		//2do parámetro = puntero en formato GO = struct donde recibe el JSON
		err = json.Unmarshal(body, &saludo)
		if err != nil {
			resJson, err := json.Marshal(responseJson{
				Message: fmt.Sprintf("Error en Unmarshal, el error es %v", err),
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

		//Unir las variables para el mensaje final
		msg := fmt.Sprintf("Hola %s %s, tienes %d años.¡Bienvenido!", saludo.Nombre, saludo.Apellido, saludo.Edad)
		//writer.Write([]byte(msg))
		//writer.WriteHeader(http.StatusAccepted)

		//Responder JSON = Marshal
		responseJson := responseNameJson{
			Message: msg,
			Status:  http.StatusOK,
		}

		res, err := json.Marshal(responseJson)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(res)
		writer.WriteHeader(http.StatusAccepted)
	})

	log.Fatal(http.ListenAndServe(":3030", nil))
}
