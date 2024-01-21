package employee

import "database/sql"

type Employee interface {
	Add(db *sql.DB)
	GetByID(db *sql.DB, id int)
	GetByName(db *sql.DB, name string)
	GetAllNames(db *sql.DB) []string
	UpdateByID(db *sql.DB, id int)
	DeleteByID(db *sql.DB, id int)
}
