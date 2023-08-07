package main

//importar los paquetes a utilizar
import (
	"encoding/json"
	"log"

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
func main() {
	//Acceder al valor que contiene el paquete
	//Inicializar valores y crear variable
	bodyRequestStructPerson := person.Person{
		Nombre: "Hola",
		Edad:   24,
	}

	bodyResponseStruct := response.Response{}

	//Marshal : Hacer la conversión de un valor a JSON, y retorna dos valores:
	// JSON en bytes y el otro un error
	jsonBodyRequestBytes, err := json.Marshal(bodyRequestStructPerson)
	if err != nil {
		//Log.Fatal : Es un log  a std out(terminal,etc) donde se
		// muestra información y después salir del programa
		log.Fatal("Error en Marshal", bodyRequestStructPerson, err)

	}

}
