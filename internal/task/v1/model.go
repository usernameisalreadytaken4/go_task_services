package task

import "time"

type Task struct {
	ID      int
	Name    string
	Created time.Time
	Updated time.Time
	Status  string
}
