package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type datos struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.POST("/validar-datos", func(c echo.Context) error {
		d := datos{}

		if err := c.Bind(&d); err != nil {
			return c.JSON(http.StatusBadRequest, response{Status: http.StatusBadRequest, Message: "body invalido"})
		}

		return c.JSON(http.StatusOK, response{Status: http.StatusOK, Message: fmt.Sprintf("Tu nombres es: %v y tu edad es %v", d.Nombre, d.Edad)})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
