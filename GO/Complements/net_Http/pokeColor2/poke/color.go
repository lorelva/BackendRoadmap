package color

type PokeColor struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Names          []Names          `json:"names"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}
type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
