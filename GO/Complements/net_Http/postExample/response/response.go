package response

//Response : Hace referencia a la Respuesta de la API
//Tags del struct  sirven para reconocer el nombre en JSON

type Response struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
