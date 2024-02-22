package loan

import (
	"database/sql"
	"log"
	"time"
)

func RequestLoan(db *sql.DB, idBook int, idUser int) {
	//FORMATO DE LA FECHA: Buscar YYYY-MM-DD HH:MM:SS
	//Crear fechas de loan date y due date con el time.now
	loanDate := time.Now().Format("2006-01-02 15:04:05")
	dueDate := time.Now().Format("2006-01-02 15:04:05")
	endDate := loanDate.AddDate(0, 0, 10)

}

func Add(db *sql.DB) {
	result, err := db.Exec(`
	INSERT INTO BOOK_LOAN (ID_BOOK, ID_USER, LOAN_DATE, DUE_DATE) VALUES (?, ?, ?, ?);	
	`)

	if err != nil {
		log.Println("Unable to insert into the LOAN table, the error is: ", err)
		return
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the aggregated columns", err)
		return
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to loan table, values are\n ID Book: %d\n ID User: %d\n Loan Date: %s\n Due Date: %s\n")
		return
	} else if rowsInserted == 0 {
		log.Println("Data unsucessfully added into the loan table", err)
		return
	}

}

func UpdateByID(db *sql.DB, id int) {
	result, err := db.Exec(`
	UPDATE BOOK_LOAN SET  ID_BOOK = ? , ID_USER = ?, LOAN_DATE = ?, DUE_DATE = ? WHERE ID = ?;
	`)
	if err != nil {
		log.Println("Data update could not be performed on the loan table, the error was:", err)
		return
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated in the loan table, the values are:")
		return
	} else if rowsUpdated == 0 {
		log.Println("Data could not be updated in the loan table")
		return
	}

}

func GetByID(db *sql.DB, id int) {
	var (
		ID        int
		ID_Book   int
		ID_User   int
		Loan_Date string
		Due_Date  string
	)

	err := db.QueryRow("SELECT * FROM BOOK_LOAN WHERE ID = ?;", id).Scan(&ID, &ID_Book, &ID_User, &Loan_Date, &Due_Date)
	if err != nil {
		log.Println("Could not retrieve information with the requested ID", err)
		return
	}

	log.Printf("Data with ID %d are: %d  %d  %s %s\n", ID, ID_Book, ID_User, Loan_Date, Due_Date)

}

func DeleteByID(db *sql.DB, id int) {
	result, err := db.Exec("DELETE FROM BOOK_LOAN WHERE ID = ?;", id)
	if err != nil {
		log.Println("Could not delete the id in the loan table, the error was: ", err)
		return
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete values with the requested id: ", err)
		return
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from the loan table", id)
		return
	} else if rowsDeleted == 0 {
		log.Println("Could not delete the ID in the loan table")
		return
	}

}
