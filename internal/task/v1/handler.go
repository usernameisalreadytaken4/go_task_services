package task

import (
	"net/http"
)

type PayloadRequest struct {
	Type TaskType
	Text string
}

type Handler struct {
	service Service
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// user := GetUserByToken()

	switch r.Method {
	case http.MethodGet:
		// h.service.GetTasksByUser(user)
	case http.MethodPost:
		// h.service.CreateTask(user, payload)
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
