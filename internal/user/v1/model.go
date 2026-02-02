package user

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	ID      int64      `json:"-"`
	User    User       `json:"-"`
	Value   string     `json:"value"`
	Created time.Time  `json:"-"`
	Updated *time.Time `json:"-"`
}

type User struct {
	ID       int64      `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"-"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated"`
}

func (u *User) SetPassword(password string) error {
	encryptedPassword, err := u.EncryptPassword(password)
	if err != nil {
		return nil
	}
	u.Password = encryptedPassword
	return nil
}

func (u *User) SetRandomPassword() error {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	return u.SetPassword(hex.EncodeToString(b))
}

func (u *User) EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	)
}
