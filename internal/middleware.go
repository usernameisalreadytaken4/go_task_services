package internal

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	userV1 "github.com/usernameisalreadytaken4/go_task_services/internal/user/v1"
)

type contextKey string

const userContextKey contextKey = "user"

func GetUserByToken(pool *pgxpool.Pool, r *http.Request) (*userV1.User, error) {
	bearer := r.Header.Get("Authorization")

	if bearer == "" || strings.HasPrefix(bearer, "Bearer ") {
		return nil, errors.New("Unauthorized")
	}

	token := strings.TrimPrefix(bearer, "Bearer ")
	hashedToken := userV1.HashToken(token)

	query := `SELECT u.id, u.email 
				FROM tokens t 
				JOIN users u ON t.user_id = u.id 
				WHERE t.value = $1`

	var user userV1.User
	err := pool.QueryRow(r.Context(), query, hashedToken).Scan(&user.ID, &user.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &user, nil
}

func AuthMiddleware(pool *pgxpool.Pool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := GetUserByToken(pool, r)
		if err != nil {
			http.Error(w, "No auth", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserFromContext(ctx context.Context) (*userV1.User, bool) {
	user, ok := ctx.Value(userContextKey).(*userV1.User)
	return user, ok
}
