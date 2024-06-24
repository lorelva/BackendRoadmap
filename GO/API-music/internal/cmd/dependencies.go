package main

import (
	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/controller"
	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/repository"
	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/server"
)

func build() {
	repo := repository.MusicRepositoryService()

	controller := controller.InitMusicController(repo)

	s := server.InitServer()
	s.Controller = controller

	s.InitMusicRoutes()

	s.StartServer()

}
