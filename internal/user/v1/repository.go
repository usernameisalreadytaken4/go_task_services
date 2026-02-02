package user

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	err := r.DB.QueryRow(
		context.Background(),
		`SELECT * FROM users WHERE email = $1`,
		user.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(email, password string) (*User, error) {

	user := &User{
		Email: email,
	}

	user.SetPassword(password)

	err := r.DB.QueryRow(
		context.Background(),
		`INSERT INTO users(email, password) VALUES ($1, $2)
		RETURNING id, email, password, created, updated`,
		user.Email, user.Password).Scan(&user.ID, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		log.Println("Insert error:", err.Error())
		return nil, ErrInternalError
	}
	return user, nil
}

func NewRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
