package queries

import (
	"database/sql"
	"log"
)

//FUnciones DML--> DATA MANUPULATION LANGUAGE == CRUD(Create, Read, Update, Delete)

func CrearTabla(db *sql.DB) {

	//Exec no retorna filas solo ejecuta comandos como: Create, Insert, Update, Delete
	_, err := db.Exec(`
	CREATE TABLE SUCURSAL(
		ID INT PRIMARY KEY ,
		DIRECCION VARCHAR(100)
	);
	`)

	if err != nil {
		//Se usa para informar de lo que sucedió , en este caso solo mandar la causa de lo
		//que pasó y no salirse del programa
		log.Println("No se pudo crear la tabla sucursal, la causa fue: ", err)
		//Sirve para salir de la función
		return
	}

	log.Println("La tabla sucursal se creó con éxito")
}

// el orden de los parametros no importa
func InsertarSucursal(db *sql.DB, direccion string, id int) {

	//como un printf en orden los parametros
	//Si quremos agregar valores propios al query se debe usar el signo: ?
	//EL orden debe ser como nuestro query lo pida o lo necesita
	//Solo aplica a mysql
	result, err := db.Exec(`INSERT INTO SUCURSAL VALUES(?, ?);`, id, direccion)
	if err != nil {
		log.Println("No se pudo insertar los datos de la tabla sucursal, la causa fue: ", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("NO se pudo obtener los valores de columnas agregadas: ", err)
		return
	}

	//Rows affected retorna un numero entero de las columnas afectadas
	//Solo aplica para insert, delete, update
	//SI rowsAffected es mayor a 0, significa que si se modificó , agregó o eliminó algún valor
	if rowsAffected > 0 {
		log.Println("Se insertó con exito en la tabla sucursal, los valores son: ", id, direccion)
		return
	} else if rowsAffected == 0 {
		log.Println("NO se pudo insertar datos en la tabla sucursal")
		return
	}

}

//para
//UPDATE FROM SUCURSAL WHERE ID = ? SET DIRECCION = ?
//DELETE FROM SUCURSAL WHERE ID = ?

func ActualizarSucursal(db *sql.DB, id int, direccion string) {

}

func EliminarSucursal() {

}
