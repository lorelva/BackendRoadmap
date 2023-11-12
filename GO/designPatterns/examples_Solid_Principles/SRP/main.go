package main

import (
	"fmt"

	"github.com/lorelva/BackendRoadmap/GO/designPatterns/examples_Solid_Principles/SRP/logger"
)

//FOLLOWING THE PRINCIPLE OF SRP

func main() {

	//SRP --> Separar el paquete por archivo(paquete logger, archivos: console.go, file.go)
	console := logger.ConsoleLogger{}
	file := logger.FileLogger{}

	//ISP --> Se seapra las funciones de los archivos(file.go, console.go) y así podemos
	//tener una segregación
	var client logger.LogClient
	msg := "hello"

	client = &console

	client.Log("")
	ok, err := client.Clean("")
	if err != nil {
		fmt.Printf("El mensaje de error es : %v ok: %v", err, ok)
	}

	client = &file
	client.Log(msg)
	ok, err = client.Clean(msg)
	if err != nil {
		fmt.Printf("El mensaje de error es : %v ok: %v", err, ok)
	}

}
