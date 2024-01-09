package worker

import (
	"database/sql"
	"log"
)

type Worker struct {
	Name string
}

func (w *Worker) Add(db *sql.DB) {
	result, err := db.Exec("INSERT INTO WORKER (NAME) VALUES(?);", w.Name)

	if err != nil {
		log.Println("No se pudo insertar los valores a la tabla Worker, el error fue :", err)
		return
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("No se pudo obtener de los valores de las columnas agregadas", err)
		return
	}

	if rowsInserted > 0 {
		log.Println("Se insertaron correctamente en la tabla worker, los valores son: ", w.Name)
		return
	} else if rowsInserted == 0 {
		log.Println("Los datos no se insertaron en la tabla worker")
		return
	}

}

func (w *Worker) GetByID(db *sql.DB, id int) {

}

func (w *Worker) GetByName(db *sql.DB, name string) {

}

func (w *Worker) GetAllNames(db *sql.DB) []string {
	//no quitar el return, por el momento con que se imprima todos los valores
	var (
		id         int
		workerName string
		names      []string
	)

	rows, err := db.Query("SELECT * FROM WORKER;")
	if err != nil {
		log.Println("No se pudo mostrar los datos de la tabla worker: ", err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&id, &workerName)
		if err != nil {
			log.Println("No se pudo obtener la información, debido a: ", err)
			return nil
		}
		//fmt.Printf("El id es: %d, el nombre  es: %s \n", id, workerName)
		names = append(names, workerName)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al recorrer las filas:", err)
		return nil
	}
	return names

}

func (w *Worker) UpdateByID(db *sql.DB, id int) {
	result, err := db.Exec("UPDATE WORKER SET NAME = ? WHERE ID = ?", w.Name, id)
	if err != nil {
		log.Println("No se pudo realizar la actualización de datos en la tabla worker, el error fue: ", err)
		return
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("NO se pudo actualizar valores de las columnas modificadas: ", err)
		return
	}

	if rowsUpdated > 0 {
		log.Println("Se actualizó con exito en la tabla worker, los valores son: ", w.Name)
		return
	} else if rowsUpdated == 0 {
		log.Println("No se pudo actualizar datos en la tabla worker")
		return
	}

}

func (w *Worker) DeleteByID(db *sql.DB, id int) {
	result, err := db.Exec("DELETE FROM WORKER WHERE ID = ?", id)
	if err != nil {
		log.Println("No se pudo eliminar el id en la tabla worker, el error fue: ", err)
		return
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("NO se pudo eliminar valores de las columnas con el id solicitado: ", err)
		return
	}

	if rowsDeleted > 0 {
		log.Printf("Se eliminó con exito el id %v en la tabla worker", id)
		return
	} else if rowsDeleted == 0 {
		log.Println("No se pudo eliminar el id en la tabla worker")
		return
	}

}
