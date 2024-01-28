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

	dataBook := book.Book{
		Title:           "loquesea3",
		Author:          "kfjsdgk",
		PublicationDate: "2027-03-10",
	}

	library.Add(&dataBook, "Christian")
	library.Add(&dataBook, "Hola")

	//dataBook.Add(libraryDB)
	//dataBook.UpdateByID(libraryDB, 3)

	//dataBook.DeleteByID(libraryDB, 4)
	//dataBook.GetByID(libraryDB, 6)

}
