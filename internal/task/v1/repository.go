package task

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	DB *pgxpool.Pool
}

type Repository interface {
	Save(*Task) (*Task, error)
	// GetByID(int) (*Task, error)
	// GetByUserID(int) ([]*Task, error)
}

func (r *repository) Save(task *Task) (*Task, error) {
	err := r.DB.QueryRow(
		context.Background(),
		`INSERT INTO tasks(user_id, payload) VALUES ($1, $2)
		RETURNING id, created`,
		task.User.ID, task.Payload).Scan(&task.ID, &task.Created)
	if err != nil {
		log.Println("Insert error:", err.Error())
		return nil, ErrInternalError
	}
	return task, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		DB: db,
	}
}
