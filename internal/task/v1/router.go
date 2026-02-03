package task

import "net/http"

func TaskRouter(mux *http.ServeMux, taskHandler *TaskHandler) {
	mux.HandleFunc("/api/v1/tasks", taskHandler.Handle)
	mux.HandleFunc("/api/v1/tasks/", taskHandler.HandleOne)
}
