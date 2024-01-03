package worker

import (
	"database/sql"
	"fmt"
)

type Worker struct {
	Name string
}

func (w *Worker) Add(db *sql.DB) {
	fmt.Println("Agregado 1")

}

func (w *Worker) GetByID(db *sql.DB, id int) {

}

func (w *Worker) GetByName(db *sql.DB, name string) {

}

func (w *Worker) GetAllNames(db *sql.DB) []string {
	//no quitar el return, por el moemnbto conq ue se imprima todos los valores
	return nil

}

func (w *Worker) UpdateByID(db *sql.DB, id int) {

}

func (w *Worker) DeleteByID(db *sql.DB, id int) {

}
