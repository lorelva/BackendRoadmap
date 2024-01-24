package librarysystem

import (
	"database/sql"

	"github.com/lorelva/designPatterns/library/librarySystem/book"
)

type Library struct {
	DB *sql.DB
}

type Management interface {
	Add(db *sql.DB)
	UpdateByID(db *sql.DB, id int)
	DeleteByID(db *sql.DB, id int)
	GetByID(db *sql.DB, id int)
}

func (l *Library) AddBook(b book.Book) {
	b.Add(l.DB)
}
