package supervisor

import (
	"database/sql"
	"log"
)

type Supervisor struct {
	Name string
}

func (s *Supervisor) Add(db *sql.DB) {
	result, err := db.Exec("INSERT INTO SUPERVISOR(NAME) VALUES(?);", s.Name)

	if err != nil {
		log.Println("No se pudo insertar los valores a la tabla Supervisor, el error fue :", err)
		return
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("No se pudo obtener de los valores de las columnas agregadas", err)
	}

	if rowsInserted > 0 {
		log.Println("Se insertaron correctamente en la tabla supervisor, los valores son: ", s.Name)
		return
	} else if rowsInserted == 0 {
		log.Println("Los datos no se insertaron en la tabla supervisor")
		return
	}
}

//Sería dificil porque se repiten id porque se tiene 1 para supervisor, 1 para worker, en cambio el nombre no importa, solo que haya repetidos de diferentes tablas
/*func (s *Supervisor) GetByID(db *sql.DB, id int) {
}
*/

func (s *Supervisor) GetByName(db *sql.DB, name string) {

}

func (s *Supervisor) GetAllNames(db *sql.DB) []string {
	var (
		id             int
		supervisorName string
		names          []string
	)

	rows, err := db.Query("SELECT * FROM SUPERVISOR;")
	if err != nil {
		log.Println("No se pudo mostrar los datos de la tabla supervisor: ", err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&id, &supervisorName)
		if err != nil {
			log.Println("No se pudo obtener la información, debido a: ", err)
			return nil
		}
		//fmt.Printf("El id es: %d, el nombre  es: %s \n", id, supervisorName)
		names = append(names, supervisorName)
	}
	return names

}

func (s *Supervisor) UpdateByID(db *sql.DB, id int) {
	result, err := db.Exec("UPDATE SUPERVISOR SET NAME = ? WHERE ID = ?", s.Name, id)
	if err != nil {
		log.Println("No se pudo realizar la actualización de datos en la tabla supervisor, el error fue: ", err)
		return
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("NO se pudo actualizar valores de las columnas modificadas: ", err)
		return
	}

	if rowsUpdated > 0 {
		log.Println("Se actualizó con exito en la tabla supervisor, los valores son: ", s.Name)
		return
	} else if rowsUpdated == 0 {
		log.Println("No se pudo actualizar datos en la tabla supervisor")
		return
	}

}

func (s *Supervisor) DeleteByID(db *sql.DB, id int) {
	result, err := db.Exec("DELETE FROM SUPERVISOR WHERE ID = ?", id)
	if err != nil {
		log.Println("No se pudo eliminar el id en la tabla supervisor, el error fue: ", err)
		return
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("No se pudo eliminar valores de las columnas con el id solicitado: ", err)
		return
	}

	if rowsDeleted > 0 {
		log.Printf("Se eliminó con exito el id %v en la tabla supervisor", id)
		return
	} else if rowsDeleted == 0 {
		log.Println("No se pudo eliminar el id en la tabla supervisor")
		return
	}

}
