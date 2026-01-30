package handlers

import "net/http"

type UserHandler struct{}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {}

func (h *UserHandler) GetToken(w http.ResponseWriter, r *http.Request)     {}
func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {}
