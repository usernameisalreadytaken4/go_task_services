package task

import (
	"encoding/json"
	"errors"
	"net/http"

	internal "github.com/usernameisalreadytaken4/go_task_services/internal"
)

type PayloadRequest struct {
	Type TaskType
	Text string
}

var (
	ErrInternalError = errors.New("Internal Error")
)

type Handler struct {
	service Service
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {

	user, ok := internal.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// h.service.GetTasksByUser(user)
	case http.MethodPost:

		var payload PayloadRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		raw, err := json.Marshal(payload.Text)
		if err != nil {
			http.Error(w, "bad payload", http.StatusBadRequest)
			return
		}

		task := &Task{
			User:    user,
			Name:    payload.Type,
			Payload: raw,
		}
		task, err = h.service.Create(user, task)
		if err != nil {
			http.Error(w, ErrInternalError.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleOne(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// h.service.GetTaskByUser(user)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
