package employee

type Supervisor struct {
	ID   int
	Name string
}

func NewSupervisor(id int, name_sup string) *Supervisor {
	return &Supervisor{
		ID:   id,
		Name: name_sup,
	}
}

func (s *Supervisor) GetSupervisor() *Supervisor {
	return s
}

func (s Supervisor) GetID() int {
	//dfsadsadadadsads
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}
