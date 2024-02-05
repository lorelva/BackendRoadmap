package student

import "database/sql"

type Student struct {
	Name      string
	Last_Name string
	Gender    string
	ID_Type   int
}

func (s *Student) Add(db *sql.DB, name string) {

}
