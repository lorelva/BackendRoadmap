package database

import (
	"database/sql"
	"log"

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

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// debemos empezar a saber usar los log de fatal cuando deseamos acabar el programa
		log.Fatalf("No se puede abrir la base de datos, el error fue %v \n", err)
	}

	// ping nos sirve para verificar la conexion
	// TODO: probar el ping en realidad
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// este valor de db contiene las funciones de mysql para ejecutar queries a la base de datos que apunta a nuestra
	// tabla TEST
	return db
}

//  sintaxis de funciones
// func saludo({paramaetros {tipo de dato}}) {valores de retorno} ...
// 1- func saludo(nombre string) (string, error) {.....}
// 2- func saludo(nombre string) string {....}
// metodos...
//  func ({variable} {nombre del struct}) saludo({parametros {tipo de dato}}) {valores de retorno} {......}
// 3- func (p Persona) saludo(nombre string) string {...}

// llamadas
// 1- valor, err := saludo("Christian") || hola, adios := "asdda", "adsdasas"
// 1.1- var valor string
//      var err error = nil
//      valor, err = saludo("Christian")
// 2- valor := saludo("Lorena")
// 2.1 - valor := ""
//       valor = saludo("Lorena")
// 3- p := Persona{}
// 3- valor := p.saludo("Lorena")
// db, err := sql.Open("mysql", cfg.FormatDSN())
// if err != nil {......}
// db.adsdadsadasdasdas

// como obtener lo que retorna una funcion
// Explicacion: de lado izquiero mis variable
// de lado derecho mi llamado a funcion
// (en este ejemplo esta funcion ficticia retorna 3 valores)
// area, lados, largo := cuadrado()
// {variables donde se guardan los valores de retorno} <- cuadrado()
// cuadrado()
