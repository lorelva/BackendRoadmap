package librarysystem

import "database/sql"

type System interface {
	Add(db *sql.DB, name string)
	UpdateByID(db *sql.DB, id int)
	DeleteByID(db *sql.DB, id int)
	GetByID(db *sql.DB, id int)
}
