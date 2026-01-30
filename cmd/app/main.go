package main

import (
	"net/http"

	v1 "github.com/usernameisalreadytaken4/go_task_services/internal/http/v1/handlers"
)

func main() {
	mux := http.NewServeMux()

	tasks := &v1.TaskHandler{}

	mux.HandleFunc("/api/v1/tasks", tasks.Handle)
	mux.HandleFunc("/api/v1/tasks/", tasks.HandleOne)
}
