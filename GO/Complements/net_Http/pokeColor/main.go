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
	pokeURL := "https://pokeapi.co/api/v2/pokemon-species/132/"

	client := &http.Client{}

	response, err := client.Get(pokeURL)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode >= 300 {
		log.Fatal(errors.New("este request dio un codigo de respuesta mayor a 300"))
	}

	defer response.Body.Close()

	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	pokes := make(map[string]interface{})

	err = json.Unmarshal(byteResponse, &pokes)
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("no puedo hacer Unmarshal de: %v", byteResponse)))
	}

	//fmt.Println(pokes)
	fmt.Println(pokes["color"])
}
