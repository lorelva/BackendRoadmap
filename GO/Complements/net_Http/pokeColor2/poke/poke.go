package poke

type PokeColor struct {
	BaseHappiness        int                 `json:"base_happiness"`
	CaptureRate          int                 `json:"capture_rate"`
	Color                Color               `json:"color"`
	EggGroups            []EggGroups         `json:"egg_groups"`
	EvolutionChain       EvolutionChain      `json:"evolution_chain"`
	EvolvesFromSpecies   any                 `json:"evolves_from_species"`
	FlavorTextEntries    []FlavorTextEntries `json:"flavor_text_entries"`
	FormDescriptions     []any               `json:"form_descriptions"`
	FormsSwitchable      bool                `json:"forms_switchable"`
	GenderRate           int                 `json:"gender_rate"`
	Genera               []Genera            `json:"genera"`
	Generation           Generation          `json:"generation"`
	GrowthRate           GrowthRate          `json:"growth_rate"`
	Habitat              Habitat             `json:"habitat"`
	HasGenderDifferences bool                `json:"has_gender_differences"`
	HatchCounter         int                 `json:"hatch_counter"`
	ID                   int                 `json:"id"`
	IsBaby               bool                `json:"is_baby"`
	IsLegendary          bool                `json:"is_legendary"`
	IsMythical           bool                `json:"is_mythical"`
	Name                 string              `json:"name"`
	Names                []Names             `json:"names"`
	Order                int                 `json:"order"`
	PalParkEncounters    []PalParkEncounters `json:"pal_park_encounters"`
	PokedexNumbers       []PokedexNumbers    `json:"pokedex_numbers"`
	Shape                Shape               `json:"shape"`
	Varieties            []Varieties         `json:"varieties"`
}
type Color struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EggGroups struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EvolutionChain struct {
	URL string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type FlavorTextEntries struct {
	FlavorText string   `json:"flavor_text"`
	Language   Language `json:"language"`
	Version    Version  `json:"version"`
}
type Genera struct {
	Genus    string   `json:"genus"`
	Language Language `json:"language"`
}
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GrowthRate struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Habitat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}
type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PalParkEncounters struct {
	Area      Area `json:"area"`
	BaseScore int  `json:"base_score"`
	Rate      int  `json:"rate"`
}
type Pokedex struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokedexNumbers struct {
	EntryNumber int     `json:"entry_number"`
	Pokedex     Pokedex `json:"pokedex"`
}
type Shape struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Varieties struct {
	IsDefault bool    `json:"is_default"`
	Pokemon   Pokemon `json:"pokemon"`
}
