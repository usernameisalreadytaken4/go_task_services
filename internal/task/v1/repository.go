package task

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}
