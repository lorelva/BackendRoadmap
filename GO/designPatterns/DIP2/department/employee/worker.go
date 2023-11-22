package employee

type Worker struct {
	ID   int
	Name string
}

func NewWorker(id int, nameWorker string) *Worker {
	return &Worker{
		ID:   id,
		Name: nameWorker,
	}

}

func (w *Worker) GetWorker() *Worker {
	return w
}

func (w Worker) GetID() int {
	return w.ID
}

func (w Worker) GetName() string {
	return w.Name
}
