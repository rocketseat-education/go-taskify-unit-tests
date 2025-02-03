package pgstore

import "github.com/jackc/pgx/v5/pgxpool"

type PGTaskStore struct {
	Queries *Queries
	Pool *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) PGTaskStore {
	return PGTaskStore{Queries: New(pool), Pool: pool}
}

func (pgtaskstore *PGTaskStore) CreateTask(title string, description string, priority int32) (store.Task, error) {
	panic("not implemented") // TODO: Implement
}
func (pgtaskstore *PGTaskStore) GetTaskById(id int32) (store.Task, error) {
	panic("not implemented") // TODO: Implement
}
func (pgtaskstore *PGTaskStore) ListTasks() ([]store.Task, error) {
	panic("not implemented") // TODO: Implement
}
func (pgtaskstore *PGTaskStore) UpdateTask(id int32, title string, description string, priority int32) (store.Task, error) {
	panic("not implemented") // TODO: Implement
}
func (pgtaskstore *PGTaskStore) DeleteTask(id int32) error {
	panic("not implemented") // TODO: Implement
}
