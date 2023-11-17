package supervisor

type Supervisor struct {
	ID   int
	Name string
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}
