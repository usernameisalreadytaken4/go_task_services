package task

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	DB *pgxpool.Pool
}

type Repository interface {
	// Save(*Task) error
	// GetByID(int) (*Task, error)
	// GetByUserID(int) ([]*Task, error)
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		DB: db,
	}
}
