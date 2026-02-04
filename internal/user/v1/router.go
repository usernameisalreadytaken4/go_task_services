package user

import "net/http"

func UserRouter(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("/api/v1/auth/register", handler.Register)
	mux.HandleFunc("/api/v1/auth/login", handler.Login)
}
