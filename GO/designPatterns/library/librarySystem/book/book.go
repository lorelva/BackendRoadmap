package book

import (
	"database/sql"
	"log"
)

type Book struct {
	Title           string
	Author          string
	PublicationDate string
}

func (b *Book) AddBook(db *sql.DB) {
	result, err := db.Exec(`
	INSERT INTO BOOK (TITLE, AUTHOR, PUBLICATION_DATE) VALUES (?, ?, ?);
	`, b.Title, b.Author, b.PublicationDate)

	if err != nil {
		log.Println("Unable to insert into the BOOK table, the error is: ", err)
		return
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the aggregated columns", err)
	}

	if rowsInserted > 0 {
		log.Println("Data successfully added into the book table, values are: ", b.Title, b.Author, b.PublicationDate)
		return
	} else if rowsInserted == 0 {
		log.Println("Data unsucessfully added into the book table", err)
		return
	}

}

func (b *Book) UpdateBook(db *sql.DB) {

}

func (b *Book) DeleteBook(db *sql.DB) {

}

func (b *Book) GetBook(db *sql.DB) {

}
