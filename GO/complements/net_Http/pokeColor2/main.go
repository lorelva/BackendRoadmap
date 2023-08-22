package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"lorelval/BackendRoadmap.com/poke"
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
		errors.New("request higher than 300")
	}

	defer response.Body.Close()

	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	pokes := poke.PokeColor{}

	// Unmarshal para pasar los bytes a un struct
	err = json.Unmarshal(byteResponse, &pokes)
	if err != nil {
		errors.New(fmt.Sprintf("Can't do Unmarshal of: %v", byteResponse))
	}

	fmt.Println(pokes.Color.Name)
	fmt.Println(pokes.Color.URL)
	fmt.Println(pokes.FlavorTextEntries[0].FlavorText)
	//fmt.Println(string(byteResponse))

}
