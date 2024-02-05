package main

import (
	"github.com/lorelva/designPatterns/library/database"
	librarysystem "github.com/lorelva/designPatterns/library/librarySystem"
	"github.com/lorelva/designPatterns/library/librarySystem/book"
	typeuser "github.com/lorelva/designPatterns/library/librarySystem/typeUser"
)

func main() {

	libraryDB := database.ConnectToLibrary()

	library := librarysystem.Library{
		DB: libraryDB,
	}

	dataBook := book.Book{
		Title:           "uvuewvueew",
		Author:          "Unknown",
		PublicationDate: "2000-01-29",
	}

	/*
		//library.Add(&System{}, "Lorena")
		library.Add()
		library.UpdateByID(&dataBook, 11)
		library.DeleteByID()
		library.GetByID()
	*/
	//dataBook.Add(libraryDB, "Lorena")
	//dataBook.UpdateByID(libraryDB, 3)
	//dataBook.DeleteByID(libraryDB, 4)
	//dataBook.GetByID(libraryDB, 6)

	user := typeuser.User{
		Name:      "Kat",
		Last_Name: "Kiro",
		Gender:    "Femenino",
		ID_Type:   1,
	}

	user.Add(libraryDB)

}
