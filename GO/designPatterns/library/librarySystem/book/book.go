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

func (b *Book) Add(db *sql.DB, name string) {
	var nameUser, typeUser string

	err := db.QueryRow(`
	SELECT NAME, UT.TYPE FROM USER
    INNER JOIN USER_TYPE UT on USER.ID_TYPE = UT.ID
    WHERE NAME = ?;
	`, name).Scan(&nameUser, &typeUser)

	if err != nil {
		log.Println("Unvailable to obtained user information:", err)
		return
	}

	if typeUser != "BIBLIOTECARIO" {
		log.Printf("User %s cannot add book,  because it's an %s user", nameUser, typeUser)
		return
	}

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
		return
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to book table, values are\n Title: %s\n Author: %s\n Publication Date: %s\n", b.Title, b.Author, b.PublicationDate)
		return
	} else if rowsInserted == 0 {
		log.Println("Data unsucessfully added into the book table", err)
		return
	}

}

func (b *Book) UpdateByID(db *sql.DB, id int, name string) {

	var nameUser, typeUser string

	err := db.QueryRow(`
	SELECT NAME, UT.TYPE FROM USER
    INNER JOIN USER_TYPE UT on USER.ID_TYPE = UT.ID
    WHERE NAME = ?;
	`, name).Scan(&nameUser, &typeUser)

	if err != nil {
		log.Println("Unvailable to obtained user information:", err)
		return
	}

	if typeUser != "BIBLIOTECARIO" {
		log.Printf("User %s cannot add book,  because it's an %s user", nameUser, typeUser)
		return
	}

	result, err := db.Exec(`
	UPDATE BOOK SET TITLE = ?, AUTHOR = ?, PUBLICATION_DATE = ? WHERE ID = ?;
	`, b.Title, b.Author, b.PublicationDate, id)
	if err != nil {
		log.Println("Data update could not be performed on the 'book' table, the error was:", err)
		return
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("NO se pudo actualizar valores de las columnas modificadas: ", err)
		return
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated in the book table, the values are:", b.Title, b.Author, b.PublicationDate)
		return
	} else if rowsUpdated == 0 {
		log.Println("Data could not be updated in the book table")
		return
	}

}

func (b *Book) DeleteByID(db *sql.DB, id int, name string) {
	var nameUser, typeUser string

	err := db.QueryRow(`
	SELECT NAME, UT.TYPE FROM USER
    INNER JOIN USER_TYPE UT on USER.ID_TYPE = UT.ID
    WHERE NAME = ?;
	`, name).Scan(&nameUser, &typeUser)

	if err != nil {
		log.Println("Unvailable to obtained user information:", err)
		return
	}

	if typeUser != "BIBLIOTECARIO" {
		log.Printf("User %s cannot add book,  because it's an %s user", nameUser, typeUser)
		return
	}

	result, err := db.Exec("DELETE FROM BOOK WHERE ID = ?;", id)
	if err != nil {
		log.Println("Could not delete the id in the 'book' table, the error was: ", err)
		return
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete values with the requested id: ", err)
		return
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from the book table", id)
		return
	} else if rowsDeleted == 0 {
		log.Println("Could not delete the ID in the book table.")
		return
	}

}

func (b *Book) GetByID(db *sql.DB, id int) {
	var (
		ID               int
		Title            string
		Author           string
		Publication_Date string
	)

	err := db.QueryRow("SELECT * FROM BOOK WHERE ID = ?;", id).Scan(&ID, &Title, &Author, &Publication_Date)
	if err != nil {
		log.Println("Could not retrieve information with the requested ID", err)
		return
	}

	log.Printf("Data with ID %d are: %s  %s  %s\n", ID, Title, Author, Publication_Date)

}
