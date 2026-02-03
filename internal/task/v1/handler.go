package task

import (
	"net/http"
)

type PayloadRequest struct {
	Type TaskType
	Text string
}

type TaskHandler struct {
	service TaskService
}

func (h *TaskHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

func (h *TaskHandler) HandleOne(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// h.service.GetTaskByUser(user)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewHandler(service *TaskService) *TaskHandler {
	return &TaskHandler{
		service: *service,
	}
}
