package server

// Esta es una función de método para la estructura ServerHandler.
// Recibe un puntero s a un objeto ServerHandler.
func (s *ServerHandler) InitMusicRoutes() *ServerHandler {
	//define una ruta en el servidor Echo
	//registra una ruta GET en el servidor Echo para manejar
	//solicitudes relacionadas con la obtención de todas las música
	s.ServerEcho.POST("/addMusic", s.Controller.CreateMusic)
	s.ServerEcho.PUT("/updateMusic", s.Controller.UpdateMusic)
	s.ServerEcho.DELETE("/deleteMusic", s.Controller.DeleteMusic)
	s.ServerEcho.GET("/getAllMusic", s.Controller.GetAllMusic)

	//Devuelve el propio objeto ServerHandler después de configurar las rutas
	return s
}
