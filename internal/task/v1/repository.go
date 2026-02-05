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
	Save(context.Context, *Task) (*Task, error)
	// GetByID(ctx, int) (*Task, error)
	// GetByUserID(ctx, int) ([]*Task, error)
}

func (r *repository) Save(ctx context.Context, task *Task) (*Task, error) {
	err := r.DB.QueryRow(
		ctx,
		`INSERT INTO tasks(user_id, payload, status) VALUES ($1, $2)
		RETURNING id, created`,
		task.User.ID, task.Payload).Scan(&task.ID, &task.Created, &task.Status)
	if err != nil {
		log.Println("Insert error:", err.Error())
		return nil, ErrInternalError
	}
	return task, nil
}

func (r *repository) GetByID(ID int) (*Task, error) {
	return nil, nil
}
func (r *repository) GetByUserID(UserID int) (*Task, error) {
	return nil, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		DB: db,
	}
}
