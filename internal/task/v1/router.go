package task

import "net/http"

func TaskRouter(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("/api/v1/tasks", handler.Handle)
	mux.HandleFunc("/api/v1/tasks/", handler.HandleOne)
}
