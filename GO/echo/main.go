package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// OFFICIAL ECHO https://echo.labstack.com/docs/quick-start
// Run the command on terminal : go get github.com/labstack/echo/v4
type jsonExample struct {
	Message string `json:"message,omitempty"`
	Age     int    `json:"age,omitempty"`
}

type jsonPerson struct {
	Name     string `json:"name,omitempty"`
	Children bool   `json:"children,omitempty"`
}

type person2 struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

type numbers struct {
	Number []int `json:"number,omitempty"`
}

func main() {
	//BASIC SYNTAX LIKE HELLO WORLD!
	//Create the server with echo.New()
	//Crear e iniciar instancia del servidor con echo.New()
	// POSTMAN : http://192.168.1.74:8080/
	e := echo.New()
	//Routing with the method GET
	//Definición de ruta por defecto  con el metódo GET
	//c echo.Context when it's exported outside -->RElated withe the request and allows to get in
	//headers, response, JSON, archivies, etc.
	e.GET("/", func(c echo.Context) error {
		//Send the message
		return c.String(http.StatusOK, "Hola, Hello World!")
	})

	//QUERY PARAMATERS WITH ? AND &
	//EXAMPLE:/show?team=x-men&member=wolverine
	// /saludo?age=12
	// /saludo equals to /saludo/
	////POSTMAN QUERY PARAMS: http://192.168.1.74:8080/saludo
	e.POST("/saludo", func(c echo.Context) error {
		json1 := jsonExample{
			Message: "Hello",
			Age:     23,
		}
		return c.JSON(http.StatusOK, json1)
	})

	//PATH PARAMETERS --> PARAMETRO DE URL  WITH /
	//EXAMPLE: /users/:id  --> be more specific on what we want directly on the path
	// /saludo/Lorena  --> Different path of  /saludo/Christian
	////POSTMAN PATH PARAMTER: http://192.168.1.74:8080/saludo/Lorena
	e.POST("/saludo/:name", func(c echo.Context) error {
		//Add the path parameter, this case : name
		name := c.Param("name")
		json1 := jsonExample{
			Message: fmt.Sprintf("Your Name is %s:", name),
			Age:     23,
		}
		return c.JSON(http.StatusOK, json1)
	})

	//QUERY PARAMETER
	// /saludo/test?secondName=Christian
	//POSTMAN QUERY PARAMS: http://192.168.1.74:8080/saludo/test?secondName=Hola
	e.POST("/saludo/test", func(c echo.Context) error {
		name := c.QueryParam("secondName")
		json1 := jsonExample{
			Message: fmt.Sprintf("Your Second Name is %s:", name),
			Age:     23,
		}
		//JSON(code int, i interface{}) error
		//interface refers to any type of data: struct, string, int, bool, is not a generic interface with global functions(methods)
		//In this case json1 = struct
		return c.JSON(http.StatusOK, json1)
	})

	//Handling Request
	//Calls the function Persona (created below)
	//POSTMAN:http://192.168.1.74:8080/saludo/persona
	e.POST("/saludo/persona", Persona)

	//STATIC CONTENT --> Servir archivos, en toda una pagina entera
	//Use pkg static where are located static archives like: .pdf, .txt, .html, images: .png, .jpg
	e.Static("/static", "static")

	e.GET("/number", Persona2)

	//http server started on [::]:8080 --> Allows any IP connections on the port that was selected
	//When there's no IP address, put the specific port :8080 == obviar el puerto en este caso :8080
	//CREATE AN INSTANCE OF LOGGER.FATAL
	e.Logger.Fatal(e.Start(":8080"))
}

// Handling Request
// RECIEVE JSON = UNMARSHALL
func Persona(c echo.Context) error {
	jsonRequest := jsonPerson{}
	//c.Bind is a like a variable to locate the union of all the BYTES of JSON into a struct
	//CONVERT BYTES into a struct in GO(UNMARSHALL)
	// PONTER & --> Recibir Dirección de memoria (donde está almacenada)
	// PONTER & --> Mandar Dirección de memoria (antes de la creacion de una variable o definición del tipo : *string)
	err := c.Bind(&jsonRequest)
	if err != nil {
		return err
	}

	// response as json with status code.
	response := jsonExample{
		Message: fmt.Sprintf("%v  %v", jsonRequest.Name, jsonRequest.Children),
	}
	return c.JSON(http.StatusOK, response)

}

func Persona2(c echo.Context) error {
	jsonRequest := person2{}
	err := c.Bind(&jsonRequest)
	if err != nil {
		return err
	}

	response := []numbers{}

	for i := 0; i < jsonRequest.Page; i++ {
		temproralNumber := []int{}
		for i := 0; i < jsonRequest.PageSize; i++ {
			temproralNumber = append(temproralNumber, i)
		}
		response = append(response, numbers{
			Number: temproralNumber,
		})
	}

	return c.JSON(http.StatusOK, response)
}
