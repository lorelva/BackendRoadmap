package librarysystem

import "database/sql"

type Management interface {
	Add(db *sql.DB, name string)
	UpdateByID(db *sql.DB, id int)
	DeleteByID(db *sql.DB, id int)
	GetByID(db *sql.DB, id int)
}
