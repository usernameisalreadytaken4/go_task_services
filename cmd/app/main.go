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

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	userRepo := userV1.NewRepository(pool)
	userService := userV1.NewService(userRepo)
	userHandler := userV1.NewHandler(userService)

	taskRepo := taskV1.NewRepository(pool)
	taskService := taskV1.NewService(taskRepo)
	taskHandler := taskV1.NewHandler(taskService)

	mux := http.NewServeMux()
	userV1.UserRouter(mux, userHandler)
	taskV1.TaskRouter(mux, taskHandler)

	log.Println("starting serve at :8080")
	http.ListenAndServe(":8080", mux)
}
