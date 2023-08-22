package person

// Person : es un struct que hace referencia al body
// Del request, es decir, lo que se va enviar
type Person struct {
	Nombre string `json:"nombre,omitempty"`
	Edad   int    `json:"edad,omitempty"`
}
