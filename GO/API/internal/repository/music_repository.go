package repository

import (
	"errors"

	"github.com/lorelva/BackendRoadmap/GO/API/internal/domain"
)

// implementación en memoria del repositorio de música
type MusicRepository struct {
	Music []domain.Music
	NewID int
}

func MusicRepositoryService() *MusicRepository {
	return &MusicRepository{
		Music: []domain.Music{},
		NewID: 1,
	}
}

// metódos
func (m *MusicRepository) Create(music *domain.Music) error {
	if music == nil {
		return errors.New("music cannot be nil")
	}
	// Crear una nueva música con el ID asignado
	addMusic := *music
	addMusic.ID = m.NewID

	// Incrementar el ID para la siguiente música
	m.NewID++

	// Añadir la nueva música a la lista
	m.Music = append(m.Music, addMusic)

	return nil
}

func (m *MusicRepository) Update(id int, musicUpdate *domain.Music) error {
	for i, music := range m.Music {
		if music.ID == id {
			// Mantener el ID original
			musicUpdate.ID = id
			// Actualizar la entrada de música
			m.Music[i] = *musicUpdate
			return nil
		}
	}
	return nil
}

func (m *MusicRepository) Delete(id int) error {
	for i, music := range m.Music {
		if music.ID == id {
			m.Music = append(m.Music[:i], m.Music[i+1:]...)
			return nil
		}
	}
	return errors.New("music with requested ID not founded")
}

func (m *MusicRepository) GetAll() ([]domain.Music, error) {
	return m.Music, nil
}
