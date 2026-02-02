package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
)

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	service UserService
}

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrInternalError     = errors.New("Internal Error")
)

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("email required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("email invalid")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) == 0 {
		return errors.New("password required")
	}
	if len(password) < 8 {
		return errors.New("password too short")
	}
	return nil
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var reg RegisterBody
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if err := ValidateEmail(reg.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := ValidatePassword(reg.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.CreateUser(reg.Email, reg.Password)
	if err != nil {
		switch err {
		case ErrInternalError:
			log.Println(err.Error())
			http.Error(w, "Internal Error. We are working on it", http.StatusInternalServerError)
		case ErrUserAlreadyExists:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			log.Println(err.Error())
			http.Error(w, "Internal Error. We are working on it", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	fmt.Println(user, pass, ok)

}

func NewHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: *service,
	}
}
