package task

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/usernameisalreadytaken4/go_task_services/internal"
)

func TaskRouter(mux *http.ServeMux, handler *Handler, db *pgxpool.Pool) {
	log.Println("test")
	fmt.Println("test")
	mux.Handle(
		"/api/v1/tasks",
		internal.AuthMiddleware(db, http.HandlerFunc(handler.Handle)),
	)

	mux.Handle(
		"/api/v1/tasks/",
		internal.AuthMiddleware(db, http.HandlerFunc(handler.HandleOne)),
	)
}
