package librarysystem

import (
	"database/sql"
)

type Library struct {
	DB *sql.DB
}

func (l *Library) Add(m Management, name string) {
	m.Add(l.DB, name)
}

//Crear las funciones correspondientes para borrar, actualizar , select
