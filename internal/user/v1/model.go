package user

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	Value string
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) SetRandomPassword() error {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	return u.SetPassword(hex.EncodeToString(b))
}
