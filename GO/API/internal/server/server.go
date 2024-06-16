package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lorelva/BackendRoadmap/GO/API/internal/controller"
)

type ServerHandler struct {
	ServerEcho *echo.Echo
	Controller *controller.MusicController
}

func InitServer() *ServerHandler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.POST, echo.HEAD, echo.PUT, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
	}))

	return &ServerHandler{
		ServerEcho: e,
	}
}

func (s *ServerHandler) StartServer() {
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
