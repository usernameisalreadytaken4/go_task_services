package user

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RandomToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func HashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

type Token struct {
	ID      int64      `json:"-"`
	UserID  int64      `json:"-"`
	Value   string     `json:"-"`
	Created time.Time  `json:"-"`
	Updated *time.Time `json:"-"`
}

func (t *Token) ValidateToken(token string) error {
	if HashToken(token) != t.Value {
		return errors.New("Invalid token")
	}
	return nil
}

func (t *Token) CreateToken() (string, error) {
	newToken := RandomToken()
	t.Value = HashToken(newToken)
	return newToken, nil
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
