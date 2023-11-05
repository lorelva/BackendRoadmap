package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	h := "06:05:03AM"

	// Layout de la hora/for the hour
	//predefined layouts for use in Time(view in documentation official of GO)
	t, err := time.Parse("03:04:05PM", h)
	if err != nil {
		log.Fatal(err)
	}

	//Output format (how you need it)
	// t.Format retorna el tiempo del tipo de dato t en el formato de la hora deseada
	horaFormateada := t.Format("15:04:05")

	fmt.Println(horaFormateada)
}
