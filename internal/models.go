package internal

import (
	"time"
)

type User struct {
	ID    int
	Login string
}

type Token struct {
	Value string
}

type Task struct {
	ID      int
	Created time.Time
	Updated time.Time
	Status  string
}
