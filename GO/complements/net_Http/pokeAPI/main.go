package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	pokeURL := "https://pokeapi.co/api/v2/pokemon/ditto"

	client := &http.Client{}

	// Get es para ejecutar el request GET hacia una URL
	response, err := client.Get(pokeURL)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode >= 300 {
		errors.New("este request dio un codigo de respuesta mayor a 300")
	}

	// defer llama a la funcion que le pongas AL FINAL de donde lo llames
	defer response.Body.Close()

	// io.ReadAll es para pasar el body (de tipo io.Reader (un archivo)) a un arreglo de bytes
	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	pokes := make(map[string]interface{})

	// Unmarshal para pasar los bytes a un mapa
	err = json.Unmarshal(byteResponse, &pokes)
	if err != nil {
		errors.New(fmt.Sprintf("no puedo hacer Unmarshal de: %v", byteResponse))
	}

	fmt.Println(pokes)
	fmt.Println(pokes["abilities"])
}
