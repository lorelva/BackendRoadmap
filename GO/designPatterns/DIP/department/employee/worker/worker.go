package worker

type Worker struct {
	ID   int
	Name string
}

func (w *Worker) GetID() int {
	return w.ID
}

func (w *Worker) GetName() string {
	return w.Name
}
