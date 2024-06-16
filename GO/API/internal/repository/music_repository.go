package repository

import (
	"errors"

	"github.com/lorelva/BackendRoadmap/GO/API/internal/domain"
)

// implementación en memoria del repositorio de música
type MusicRepository struct {
	music  []domain.Music
	nextID int
}

func MusicRepositoryService() *MusicRepository {
	return &MusicRepository{
		music:  []domain.Music{},
		nextID: 1,
	}
}

// metódos
func (m *MusicRepository) Create(music *domain.Music) error {
	music.ID = m.nextID
	m.nextID++
	m.music = append(m.music, *music)
	return nil
}

func (m *MusicRepository) Update(id int, musicUpdate *domain.Music) error {
	for i, music := range m.music {
		if music.ID == id {
			musicUpdate.ID = id
			m.music[i] = *musicUpdate
			return nil
		}
	}
	return nil
}

func (m *MusicRepository) Delete(id int) error {
	for i, music := range m.music {
		if music.ID == id {
			m.music = append(m.music[:i], m.music[i+1:]...)
			return nil
		}
	}
	return errors.New("music with requested ID not founded")
}

func (m *MusicRepository) GetAllOrByID(id int) ([]domain.Music, error) {
	musicList := []domain.Music{}
	if id > 0 {
		for _, music := range m.music {
			if music.ID == id {
				musicList = append(musicList, music)
				break
			}
		}
		if len(musicList) == 0 {
			return []domain.Music{}, errors.New("music with requested ID not found")
		}
	} else {
		musicList = m.music
	}

	return musicList, nil
}
