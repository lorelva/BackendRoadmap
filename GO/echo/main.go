package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jsonExample struct {
	Message string `json:"message,omitempty"`
	Age     int    `json:"age,omitempty"`
}

type jsonPerson struct {
	Name     string `json:"name,omitempty"`
	Children bool   `json:"children,omitempty"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola")
	})

	//QUERY PARAMATER
	// /saludo?age=12
	e.POST("/saludo", func(c echo.Context) error {
		json1 := jsonExample{
			Message: "Hello",
			Age:     23,
		}
		return c.JSON(http.StatusOK, json1)
	})

	//PARAMETER
	// /saludo/Lorena
	// /saludo/Christian
	e.POST("/saludo/:name", func(c echo.Context) error {
		name := c.Param("name")
		json1 := jsonExample{
			Message: fmt.Sprintf("Your Name is %s:", name),
			Age:     23,
		}
		return c.JSON(http.StatusOK, json1)
	})

	//QUERY PARAMETER
	// /saludo/test?secondName=Christian
	e.POST("/saludo/test", func(c echo.Context) error {
		name := c.QueryParam("secondName")
		json1 := jsonExample{
			Message: fmt.Sprintf("Your Second Name is %s:", name),
			Age:     23,
		}
		return c.JSON(http.StatusOK, json1)
	})

	e.POST("/saludo/persona", Persona)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":8080"))
}

func Persona(c echo.Context) error {
	jsonRequest := jsonPerson{}
	err := c.Bind(&jsonRequest)
	if err != nil {
		return err
	}

	response := jsonExample{
		Message: fmt.Sprintf("%v  %v", jsonRequest.Name, jsonRequest.Children),
	}
	return c.JSON(http.StatusOK, response)

}
