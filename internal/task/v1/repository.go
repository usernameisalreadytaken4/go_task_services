package task

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	DB *pgxpool.Pool
}

type Repository interface {
	Save(context.Context, *Task) (*Task, error)
	GetByID(context.Context, int, int) (*Task, error)
	GetByUserID(context.Context, int) ([]Task, error)
}

func (r *repository) Save(ctx context.Context, task *Task) (*Task, error) {
	err := r.DB.QueryRow(
		ctx,
		`INSERT INTO tasks(user_id, payload) VALUES ($1, $2)
		RETURNING id, created, status`,
		task.User.ID, task.Payload).Scan(&task.ID, &task.Created, &task.Status)
	if err != nil {
		log.Println("Insert error:", err.Error())
		return nil, ErrInternalError
	}
	return task, nil
}

func (r *repository) GetByID(ctx context.Context, ID, userID int) (*Task, error) {
	var task Task

	err := r.DB.QueryRow(
		ctx,
		`SELECT id, type, created, started, finished, status, payload, result 
		FROM tasks
		WHERE id = $1 and user_id = $2`,
		ID, userID).Scan(&task.ID, &task.Type, &task.Created, &task.Started, &task.Finished, &task.Status, &task.Payload, &task.Result)

	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		log.Println(err.Error())
		return nil, ErrInternalError
	}
	return &task, nil
}

func (r *repository) GetByUserID(ctx context.Context, userID int) ([]Task, error) {
	var tasks []Task

	rows, err := r.DB.Query(
		ctx,
		`SELECT id, type, created, started, finished, status, payload, result 
		FROM tasks
		WHERE user_id = $1`,
		userID)
	if err != nil {
		log.Println(err.Error())
		return nil, ErrInternalError
	}

	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Type,
			&task.Created,
			&task.Started,
			&task.Finished,
			&task.Status,
			&task.Payload,
			&task.Result,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, ErrInternalError
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		DB: db,
	}
}
