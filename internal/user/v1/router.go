package user

import "net/http"

func UserRouter(mux *http.ServeMux, userHandler *UserHandler) {
	mux.HandleFunc("/api/v1/auth/register", userHandler.Register)
	mux.HandleFunc("/api/v1/auth/login", userHandler.Login)
}
