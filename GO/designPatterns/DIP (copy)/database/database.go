package database

import (
	"database/sql"
	"log"

	//IMPORT ANONIMO
	//En caso de no el paquete de mysql se debe importar agregando el guion bajo al inicio
	//y no se borre
	//EJEMPLO: _ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
)

// ConnectionDB retorna un puntero de sql.DB con la base de datos conectada a la que creamos,
// en este caso especifico nos va a retornar la informacion de la tabla TEST
func ConnectToTestTable() *sql.DB {
	cfg := mysql.Config{
		User:   "lorelva",
		Passwd: "140218",
		Addr:   "localhost:3306",
		DBName: "TEST",
		Net:    "tcp",
	}

	//Tiene como parametro sql.Open el strind de driver y el segundp el string de conexión
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// debemos empezar a saber usar los log de fatal cuando deseamos acabar el programa
		log.Fatalf("No se puede abrir la base de datos, el error fue: %v \n", err)
	}

	//El ping siempre debe ser usado porque si no se aprecen bugs o errores.
	// ping nos sirve para verificar la conexion
	// TODO: probar el ping en realidad

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo hacer ping , el error fue :", err)
	}

	// este valor de db contiene las funciones de mysql para ejecutar queries a la base de datos que apunta a nuestra
	// tabla TEST
	//retorno del puntero de sql
	return db
}

func ConnectToDepartment() *sql.DB {
	cfg := mysql.Config{
		User:   "lorelva",
		Passwd: "140218",
		Addr:   "localhost:3306",
		DBName: "DEPARTMENT",
		Net:    "tcp",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("No se pudo abrir la base de datos, debido a: %v \n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo realizar ping con la base de datos, el eeror fue:", err)
	}
	return db

}
