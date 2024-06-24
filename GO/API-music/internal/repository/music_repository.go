package repository

import (
	"errors"

	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/domain"
)

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

func (m *MusicRepository) Create(music *domain.Music) error {
	if music == nil {
		return errors.New("music cannot be nil")
	}

	addMusic := *music
	addMusic.ID = m.NewID
	m.NewID++

	m.Music = append(m.Music, addMusic)
	return nil
}

func (m *MusicRepository) Update(id int, musicUpdate *domain.Music) error {
	for i, music := range m.Music {
		if music.ID == id {
			musicUpdate.ID = id
			m.Music[i] = *musicUpdate
			return nil
		}
	}

	return errors.New("music doesn't exist")
}

func (m *MusicRepository) Delete(id int) error {
	for i, music := range m.Music {
		if music.ID == id {
			before := m.Music[:i]
			after := m.Music[i+1:]

			m.Music = append(before, after...)
			return nil

		}
	}

	return errors.New("music with requested ID not founded")
}

func (m *MusicRepository) Get() ([]domain.Music, error) {
	return m.Music, nil
}
