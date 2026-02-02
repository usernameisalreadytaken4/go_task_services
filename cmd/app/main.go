package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
	userV1 "github.com/usernameisalreadytaken4/go_task_services/internal/user/v1"
)

func main() {

	dsn := "postgres://postgres:my_password@localhost:5432/task_service" // забирать из конфига

	config, _ := pgxpool.ParseConfig(dsn)
	pool, _ := pgxpool.NewWithConfig(context.Background(), config)

	err := pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	userRepo := userV1.NewRepository(pool)
	userService := userV1.NewService(userRepo)
	userHandler := userV1.NewHandler(userService)

	tasks := &taskV1.TaskHandler{
		DB: pool,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/auth/register", userHandler.Register)
	mux.HandleFunc("/api/v1/auth/login", userHandler.Login)

	mux.HandleFunc("/api/v1/tasks", tasks.Handle)
	mux.HandleFunc("/api/v1/tasks/", tasks.HandleOne)

	http.Handle("/", mux)

	log.Println("starting serve at :8080")
	http.ListenAndServe(":8080", nil)
}
