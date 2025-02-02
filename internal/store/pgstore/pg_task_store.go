package pgstore

import "github.com/jackc/pgx/v5/pgxpool"

type PGTaskStore struct {
	Queries *Queries
	Pool *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) PGTaskStore {
	return PGTaskStore{Queries: New(pool), Pool: pool}
}
