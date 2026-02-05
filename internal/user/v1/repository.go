package user

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
	Create(context.Context, string, string) (*User, error)
	GetByEmail(context.Context, string) (*User, error)
	GetToken(context.Context, User) (string, error)
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.DB.QueryRow(
		ctx,
		`SELECT * FROM users WHERE email = $1`,
		email).Scan(&user.ID, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(ctx context.Context, email, password string) (*User, error) {

	user := &User{
		Email: email,
	}

	user.SetPassword(password)

	err := r.DB.QueryRow(
		ctx,
		`INSERT INTO users(email, password) VALUES ($1, $2)
		RETURNING id, email, password, created, updated`,
		user.Email, user.Password).Scan(&user.ID, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		log.Println("Insert error:", err.Error())
		return nil, ErrInternalError
	}
	return user, nil
}

func (r *repository) GetToken(ctx context.Context, user User) (string, error) {
	var token Token

	err := r.DB.QueryRow(
		ctx,
		`SELECT value FROM tokens WHERE user_id = $1`,
		user.ID).Scan(&token.Value)
	if err != nil && err != pgx.ErrNoRows {
		return "", err
	}

	// добавить транзакцию на случай обрывания контекста
	if err == nil {
		r.DB.Exec(ctx, `DELETE FROM tokens WHERE user_id = $1`, user.ID)
	}

	tokenName, _ := token.CreateToken()
	r.DB.Exec(
		ctx,
		`INSERT INTO tokens (user_id, value) VALUES ($1, $2)`,
		user.ID, token.Value)

	return tokenName, nil

}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		DB: db,
	}
}
