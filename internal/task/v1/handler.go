package task

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	internal "github.com/usernameisalreadytaken4/go_task_services/internal"
)

type PayloadRequest struct {
	Type TaskType
	Text string
}

var (
	ErrInternalError = errors.New("Internal Error")
	ErrNotFound      = errors.New("Not found")
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
		task, err := h.service.GetByUserID(r.Context(), user.ID)
		if err != nil {
			switch err {
			case ErrNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			default:
				log.Println(err)
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(task)
		return
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
			Type:    payload.Type,
			Payload: raw,
		}
		task, err = h.service.Create(r.Context(), user, task)
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

	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]

	ID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	user, ok := internal.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, err := h.service.Get(r.Context(), ID, user.ID)
		if err != nil {
			switch err {
			case ErrNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			default:
				log.Println(err)
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(task)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
