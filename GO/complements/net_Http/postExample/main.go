package main

//importar los paquetes a utilizar, en este caso de person y response
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lorelva/BackendRoadmap/person"
	"github.com/lorelva/BackendRoadmap/response"
)

// Pasos a seguir
// 1.- Definir y declarar las estructuras necesarias
// 2.- Hacer la conversión de bodyRequestStructPerson a JSON porque la API espera un JSON
// 3.- Se procede a realizar el request: POST : localhost:1323/validar-datos
// 4.- Se lee la respuesta de la API
// 5.- Se realiza la conversión del body de la respeusta en JSON a struct : bodyResponseStruct
// 6.- Validar respuesta y mostrar información

// PASO 1
func main() {
	//Acceder al valor que contiene el paquete
	//Inicializar valores y crear variable
	bodyRequestStructPerson := person.Person{
		Nombre: "Lorena",
		Edad:   24,
	}

	bodyResponseStruct := response.Response{}

	//PASO 2
	//Marshal : Hacer la conversión de un valor a JSON, y retorna dos valores:
	// JSON en bytes y el otro un error
	jsonBodyRequestBytes, err := json.Marshal(bodyRequestStructPerson)
	if err != nil {
		//Log.Fatal : Es un log  a std out(terminal,etc) donde se
		// muestra información y después se sale del programa
		log.Fatal("Error en Marshal", bodyRequestStructPerson, err)

	}

	//Se crea el cliente http: permite hacer peticiones con configuraciones especificas
	client := &http.Client{}

	//PASO 3
	//Llamar el metodo post hacia la url , content type y el paquete de bytes que se van a utilizar
	response, err := client.Post(
		"localhost:1323/validar-datos",
		"application/json",
		bytes.NewBuffer(jsonBodyRequestBytes),
	)

	//Defer se usa para ejecutarlo hasta el final del response y cerrarlo
	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("La respuesta fue exitosa")
	}

	if response.StatusCode >= 300 {
		log.Fatal(errors.New("La respuesta no se realizó con éxito, hay un error"))
	}

	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteResponse, &bodyRequestStructPerson)
	if err != nil {
		errors.New(fmt.Sprintf("No se peude hacer marshal de:%v", byteResponse))

	}

}
