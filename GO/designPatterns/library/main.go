package main

import (
	"github.com/lorelva/designPatterns/library/database"
	librarysystem "github.com/lorelva/designPatterns/library/librarySystem"
	"github.com/lorelva/designPatterns/library/librarySystem/book"
)

func main() {

	libraryDB := database.ConnectToLibrary()

	library := librarysystem.Library{
		DB: libraryDB,
	}

	b := book.Book{
		Title:           "Hola",
		Author:          "lorena",
		PublicationDate: "2000-07-23",
	}

	b.AddBook(libraryDB)

}
