package queries

import (
	"database/sql"
	"fmt"
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
		log.Println("NO se pudo obtener los valores de las columnas agregadas: ", err)
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

func ActualizarSucursal(db *sql.DB, direccion string, id int) {
	result, err := db.Exec("UPDATE SUCURSAL SET DIRECCION = ? WHERE ID = ?", direccion, id)
	if err != nil {
		log.Println("No se pudo realizar la actualización de datos en la tabla sucursal, el error fue: ", err)
		return
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("NO se pudo actualizar valores de las columnas modificadas: ", err)
		return
	}

	if rowsUpdated > 0 {
		log.Println("Se actualizó con exito en la tabla sucursal, los valores son: ", id, direccion)
		return
	} else if rowsUpdated == 0 {
		log.Println("NO se pudo actualizar datos en la tabla sucursal")
		return
	}
}

func EliminarSucursal(db *sql.DB, id int) {
	result, err := db.Exec("DELETE FROM SUCURSAL WHERE ID = ?", id)
	if err != nil {
		log.Println("No se pudo eliminar el id en la tabla sucursal, el error fue: ", err)
		return
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("NO se pudo eliminar valores de las columnas con el id solicitado: ", err)
		return
	}

	if rowsDeleted > 0 {
		log.Printf("Se eliminó con exito el id %v en la tabla sucursal", id)
		return
	} else if rowsDeleted == 0 {
		log.Println("NO se pudo eliminar el id en la tabla sucursal")
		return
	}
}

func ObtenerSucursal(db *sql.DB) {
	var (
		id        int
		direccion string
	)

	rows, err := db.Query("SELECT * FROM SUCURSAL;")
	if err != nil {
		log.Println("No se pudo mostrar los datos de la tabla sucursal: ", err)
		return
	}

	//en go no hay do while or while
	//en este caso se usa un for para recorrer cada valor de la tabla(es como si fuese un while)
	for rows.Next() {
		//Siempre se lee de forma ordenda conforme a la tabla
		//Se debe pasar la direccion de mmemoria de neutsra variable a guardar
		err = rows.Scan(&id, &direccion)
		//si error  no es nulo-->tiene algo
		//si es nil -->no hay nada
		if err != nil {
			log.Println("No se pudo obtener la información, debido a: ", err)
			return
		}
		fmt.Printf("El id es: %d, la direccion es: %s \n", id, direccion)
	}
}

func ObtenerSucursal2(db *sql.DB) {
	var (
		id int
		//Cuando hay valores nulos en la tabla , se deben usar los tipos de datos
		//Del paquete database/sql de GO: NullSTring, NUllInt.....
		direccion sql.NullString
	)

	rows, err := db.Query("SELECT * FROM SUCURSAL;")
	if err != nil {
		log.Println("No se pudo mostrar los datos de la tabla sucursal: ", err)
		return
	}

	//en go no hay do while or while
	//en este caso se usa un for para recorrer cada valor de la tabla(es como si fuese un while)
	for rows.Next() {
		//Siempre se lee de forma ordenda conforme a la tabla
		//Se debe pasar la direccion de mmemoria de neutsra variable a guardar
		err = rows.Scan(&id, &direccion)
		//si error  no es nulo-->tiene algo
		//si es nil -->no hay nada
		if err != nil {
			log.Println("No se pudo obtener la información, debido a: ", err)
			return
		}
		fmt.Printf("El id es: %d, la direccion es: %s \n", id, direccion.String)
	}
}

func ObtenerSucursalPorID(db *sql.DB, idSucursal int) {
	var (
		id int
	)

	rows, err := db.Query("SELECT ID FROM SUCURSAL WHERE ID = ?;")
	if err != nil {
		log.Println("No se pudo mostrar los datos de la tabla sucursal: ", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Println("No se pudo obtener la información, debido a: ", err)
			return
		}
		fmt.Printf("El id es: %d", id)
	}
}

//QUERIES PARA LA BASE DE DATOS DEPARTMENT EN CASO DE QUE SE QUIEREN CREAR DESDE CODIGO Y NO EN LA BASE DE DATOS
//WORKER Y SUPERVISOR
/*
func CreateTableWorker(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE WORKER(
		ID INT PRIMARY KEY AUTO_INCREMENT,
		NAME VARCHAR(100)
	);
	`)

	if err != nil {
		log.Println("No se pudo crear la tabla sucursal, la causa fue: ", err)
		return
	}

	log.Println("La tabla worker se creó con éxito")
}

func CreateTableSupervisor(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE SUPERVISOR(
		ID INT PRIMARY KEY AUTO_INCREMENT,
		NAME VARCHAR(100)
	);
	`)

	if err != nil {
		log.Println("No se pudo crear correctamente la tabla supervisor, la causa fue: ", err)
		return
	}
	log.Println("La tabla worker se creó con éxito")

}
*/
