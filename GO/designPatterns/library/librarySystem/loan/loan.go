package loan

import (
	"database/sql"
	"time"
)

func RequestLoan(db *sql.DB, idBook int, idUser int) {
	//FORMATO DE LA FECHA: Buscar YYYY-MM-DD HH:MM:SS
	//Crear fechas de loan date y due date con el time.now

	loanDate := time.Now()
	dueDate := time.Now()

}
