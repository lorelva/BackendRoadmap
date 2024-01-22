package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectToLibrary() *sql.DB {
	cfg := mysql.Config{
		User:   "lorelva",
		Passwd: "140218",
		Addr:   "localhost",
		DBName: "LIBRARY",
		Net:    "tcp",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Cannot connect to Database, error founded: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Couldn't ping the database, the error was:", err)
	}

	return db

}
