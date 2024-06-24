package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/controller"
)

type ServerHandler struct {
	ServerEcho *echo.Echo
	//controlador específico para manejar operaciones relacionadas con la música dentro del servidor.
	//Separar la lógica del controlador de la lógica del servidor
	Controller *controller.MusicController
}

// inicializa y configura el servidor Echo
func InitServer() *ServerHandler {
	//Aquí se crea una nueva instancia de Echo.
	//Esto inicializa un servidor HTTP básico que se puede configurar y ejecutar.
	e := echo.New()

	//para añadir funcionalidades comunes al servidor como registro, recuperación de errores y gestión de CORS.
	//Registra cada solicitud HTTP que llega al servidor
	e.Use(middleware.Logger())
	//// Recupera el control en caso de que ocurra un error durante el manejo de solicitudes
	e.Use(middleware.Recover())
	//Configura las políticas de CORS (Cross-Origin Resource Sharing).
	//Específicamente, permite que las solicitudes de http://localhost:3000
	//especifica los encabezados/cabeceras HTTP permitidos (Origin, Content-Type, Accept)
	//métodos HTTP permitidos
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
	}))

	//// Asigna la instancia de Echo configurada al ServerHandler
	return &ServerHandler{
		ServerEcho: e,
	}
}

// Inicia el servidor Echo en el puerto :3000
func (s *ServerHandler) StartServer() {
	//Si ocurre un error al iniciar el servidor, este método imprime un mensaje de error
	//y finaliza la ejecución del programa.
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
