package repository

import "github.com/lorelva/BackendRoadmap/GO/API-music/internal/domain"

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
	m.NewID++
	music.ID = m.NewID
	m.Music = append(m.Music, *music)
	return nil
}

func (m *MusicRepository) Update(music *domain.Music) error {
	return nil
}

func (m *MusicRepository) Delete(id int) error {
	return nil
}

func (m *MusicRepository) Get() ([]domain.Music, error) {

	return nil, nil
}
