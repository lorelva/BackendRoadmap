package main

import (
	"fmt"

	"github.com/lorelva/designPatterns/library/database"
	librarysystem "github.com/lorelva/designPatterns/library/librarySystem"
	"github.com/lorelva/designPatterns/library/librarySystem/book"
)

func main() {

	libraryDB := database.ConnectToLibrary()

	l := librarysystem.Library{
		DB: libraryDB,
	}

	dataBook := book.Book{
		Title:           "jiji",
		Author:          "uiuu",
		PublicationDate: "2023-03-10",
	}

	//l.AddBook(dataBook)
	fmt.Println(l)

	//dataBook.Add(libraryDB)
	//dataBook.UpdateByID(libraryDB, 3)

	dataBook.DeleteByID(libraryDB, 4)
	//dataBook.GetByID(libraryDB, 6)

}
