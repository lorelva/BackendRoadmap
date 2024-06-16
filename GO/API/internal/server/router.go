package server

func (s *ServerHandler) InitMusicRoutes() *ServerHandler {
	s.ServerEcho.POST("/addMusic", s.Controller.CreateMusic)
	s.ServerEcho.PUT("/updateMusic", s.Controller.UpdateMusic)
	s.ServerEcho.DELETE("/deleteMusic", s.Controller.DeleteMusic)
	s.ServerEcho.GET("/getAllMusic", s.Controller.GetAllMusic)

	return s
}
