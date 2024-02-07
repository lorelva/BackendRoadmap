package main

import (
	"github.com/lorelva/designPatterns/library/database"
	librarysystem "github.com/lorelva/designPatterns/library/librarySystem"
	"github.com/lorelva/designPatterns/library/librarySystem/book"
	"github.com/lorelva/designPatterns/library/librarySystem/loan"
	typeuser "github.com/lorelva/designPatterns/library/librarySystem/typeUser"
)

func main() {

	libraryDB := database.ConnectToLibrary()

	library := librarysystem.Library{
		DB: libraryDB,
	}

	dataBook := book.Book{}

	//library.Add(&System{}, "Lorena")
	//library.Add(&dataBook, "L")
	//MANDAR LA FECHA EN FORMATO YYYY-MM-DD EN LA LINEA 33 --> sustituir por el time.now

	/*
		format := time.Now().Format("2006-01-02")
		library.UpdateByID(&book.Book{
			Title:           "actualzizado",
			Author:          "fbsbf",
			PublicationDate: format,
		}, 11, "Lorena")
	*/
	library.GetByID(&dataBook, 11)
	library.DeleteByID(&dataBook, 11, "Lorena")
	library.GetByID(&dataBook, 11)

	//dataBook.Add(libraryDB, "Lorena")
	//dataBook.UpdateByID(libraryDB, 3)
	//dataBook.DeleteByID(libraryDB, 4)
	//dataBook.GetByID(libraryDB, 6)

	user := typeuser.User{
		Name:     "Becky",
		LastName: "hayak",
		Gender:   "Femenino",
		IDType:   1,
	}

	//user.Add(libraryDB)
	//user.UpdateByID(libraryDB, 3)
	user.GetByID(libraryDB, 4)
	//user.DeleteByID(libraryDB, 3)

	loan.RequestLoan()

}
