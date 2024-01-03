package supervisor

import (
	"database/sql"
	"fmt"
)

type Supervisor struct {
	Name string
}

func (s *Supervisor) Add(db *sql.DB) {
	fmt.Println("Agregado 1")

}

func (s *Supervisor) GetByID(db *sql.DB, id int) {

}

func (s *Supervisor) GetByName(db *sql.DB, name string) {

}

func (s *Supervisor) GetAllNames(db *sql.DB) []string {
	return nil

}

func (s *Supervisor) UpdateByID(db *sql.DB, id int) {

}

func (s *Supervisor) DeleteByID(db *sql.DB, id int) {

}
