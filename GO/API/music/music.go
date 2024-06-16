package music

import (
	"github.com/lorelva/BackendRoadmap/GO/API/internal/domain"
)

type MusicRepository interface {
	Create(music *domain.Music) error
	Update(id int, musicUpdate *domain.Music) error
	Delete(id int) error
	GetAll() ([]domain.Music, error)
	GetByID(id int) (*domain.Music, error)
}
