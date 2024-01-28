package librarysystem

import (
	"database/sql"
)

type Library struct {
	DB *sql.DB
}

func (l *Library) Add(s System, name string) {
	s.Add(l.DB, name)
}

// Crear las funciones correspondientes para borrar, actualizar , select
func (l *Library) UpdateByID(s System, id int) {
	s.UpdateByID(l.DB, id)
}

func (l *Library) DeleteByID(s System, id int) {
	s.DeleteByID(l.DB, id)
}

func (l *Library) GetByID(s System, id int) {
	s.GetByID(l.DB, id)
}
