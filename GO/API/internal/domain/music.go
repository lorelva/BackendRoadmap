package domain

// Music representa la entidad de datos para la música
type Music struct {
	ID           int    `json:"id,omitempty"`
	SongTitle    string `json:"song_title,omitempty"`
	Artist       string `json:"artist,omitempty"`
	Album        string `json:"album,omitempty"`
	MusicalGenre string `json:"musical_genre,omitempty"`
}
