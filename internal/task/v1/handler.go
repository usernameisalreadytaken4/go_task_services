package task

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskHandler struct {
	DB *pgxpool.Pool
}

func (h *TaskHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("METHOD GET")
	case http.MethodPost:
		log.Println("METHOD POST")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) HandleOne(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Method GET")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
