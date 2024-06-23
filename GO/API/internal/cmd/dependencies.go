package main

import (
	"github.com/lorelva/BackendRoadmap/GO/API/internal/controller"
	"github.com/lorelva/BackendRoadmap/GO/API/internal/repository"
	"github.com/lorelva/BackendRoadmap/GO/API/internal/server"
)

// Esta función es el punto de entrada para construir y arrancar el servidor
func build() {
	//Aquí se inicializa el repositorio
	repo := repository.MusicRepositoryService()
	//e inicializa el controlador de música
	controller := controller.InitMusicController(repo)

	//se inicializa el servidor Echo. InitServer crea una instancia de ServerHandler
	s := server.InitServer()
	//Asignación del Controlador al Servidor
	s.Controller = controller
	// para definir las rutas HTTP que manejarán las solicitudes relacionadas con la música.
	//Las rutas están asociadas con las funciones del controlador que procesan las solicitudes HTTP.
	s.InitMusicRoutes()

	//del servidor para comenzar a escuchar en el puerto especificado (por defecto en el puerto 3000).
	s.StartServer()
}
