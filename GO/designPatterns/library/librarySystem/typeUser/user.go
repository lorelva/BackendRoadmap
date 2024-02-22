package typeuser

import (
	"database/sql"
	"log"
)

type User struct {
	Name     string
	LastName string
	Gender   string
	IDType   int
}

func (u *User) Add(db *sql.DB, name string) {
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
	INSERT INTO USER (NAME, LAST_NAME, GENDER, ID_TYPE)
    VALUES (?, ?, ?, ?);
	`, u.Name, u.LastName, u.Gender, u.IDType)

	if err != nil {
		log.Println("Unable to insert into the user table, the error is: ", err)
		return
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the aggregated columns", err)
		return
	}

	if rowsInserted > 0 {
		log.Printf(
			"Data successfully added to user table, values are\n Name: %s\n Last Name: %s\n Gender: %s\n, ID Type: %d\n",
			u.Name, u.LastName, u.Gender, u.IDType)
		return
	} else if rowsInserted == 0 {
		log.Println("Data unsucessfully added into  user table", err)
		return
	}
}

func (u *User) UpdateByID(db *sql.DB, id int, name string) {
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
	UPDATE USER SET NAME= ?, LAST_NAME = ?, GENDER= ?, ID_TYPE = ? WHERE ID= ?;
	`, u.Name, u.LastName, u.Gender, u.IDType, id)

	if err != nil {
		log.Println("Data update could not be on user table, the error was:", err)
		return
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return
	}

	if rowsUpdated > 0 {
		log.Println("User table was successfully updated , the values are:", u.Name, u.LastName, u.Gender, u.IDType)
		return
	} else if rowsUpdated == 0 {
		log.Println("Data could not be updated in the book table")
		return
	}
}

func (u *User) DeleteByID(db *sql.DB, id int, name string) {
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

	result, err := db.Exec("DELETE FROM USER WHERE ID = ?;", id)
	if err != nil {
		log.Println("Could not delete the id on the user table, the error was: ", err)
		return
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete values with the requested ID: ", err)
		return
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from the user table", id)
		return
	} else if rowsDeleted == 0 {
		log.Println("Could not delete the ID in the user table")
		return
	}

}

func (u *User) GetByID(db *sql.DB, id int) {
	var (
		ID       int
		Name     string
		LastName string
		Gender   string
		Id_type  int
	)

	err := db.QueryRow("SELECT * FROM USER WHERE ID = ?;", id).Scan(&ID, &Name, &LastName, &Gender, &Id_type)
	if err != nil {
		log.Println("Information with the requested ID,doesn't exists", err)
		return
	}
	log.Printf("Data with ID %d are: %s  %s  %s %d\n", ID, Name, LastName, Gender, Id_type)

}
