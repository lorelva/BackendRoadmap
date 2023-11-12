package logger

// Principio de interfaz porque segregamos los metodos de FileLogger y de ConsoleLogger.
// Principio de susticion de liskov porque podemos modificar nuestros metodos
// (SIN CAMBIAR LA ESTRUCTURA DE LA FUNCION(EJ, NOMBRE, PARAMETROS, VALORES DE RETORNO))
type LogClient interface {
	Log(message string)
	Clean(msg string) (bool, error)
}
