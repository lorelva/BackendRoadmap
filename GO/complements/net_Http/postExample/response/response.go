package response

//Response : Hace referencia a un struct de la Respuesta de la API
//ES decir, lo que va a devolver
//Tags del struct  sirven para reconocer el nombre en JSON

type Response struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
