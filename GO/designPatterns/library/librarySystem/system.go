package librarysystem

import "database/sql"

type System interface {
	Add(db *sql.DB, name string)
	UpdateByID(db *sql.DB, id int, name string)
	DeleteByID(db *sql.DB, id int, name string)
	GetByID(db *sql.DB, id int)
}
