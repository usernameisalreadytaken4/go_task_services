package user

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/mail"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	DB *pgxpool.Pool
}

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
	user, err := h.createUser(reg.Email, reg.Password)
	if err != nil {
		log.Println("DB error:", err.Error()) // логируем и делаем отправку в центри, например
		http.Error(w, "Internal Error. We are working on it", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func (h *UserHandler) createUser(email, password string) (*User, error) {
	user := &User{
		Email: email,
	}
	err := h.DB.QueryRow(
		context.Background(),
		`SELECT * FROM users WHERE email = $1`,
		user.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil && err != pgx.ErrNoRows {
		log.Println(err.Error())
		return user, errors.New("Internal Error, we working on at")
	}
	if err != pgx.ErrNoRows {
		// придумать отдельные статус коды
		return user, errors.New("User is exists")
	}

	user.SetPassword(password)

	err = h.DB.QueryRow(
		context.Background(),
		`INSERT INTO users(email, password) VALUES ($1, $2)
		RETURNING id, email, password, created, updated`,
		user.Email, user.Password).Scan(&user.ID, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}
