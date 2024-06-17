package music

import (
	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/domain"
)

type Music interface {
	Create(music *domain.Music) error
	Update(music *domain.Music) error
	Delete(id int) error
	Get() ([]domain.Music, error)
}
