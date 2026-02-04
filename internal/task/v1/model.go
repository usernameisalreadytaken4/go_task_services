package task

import (
	"encoding/json"
	"time"

	userV1 "github.com/usernameisalreadytaken4/go_task_services/internal/user/v1"
)

type TaskType string

type Task struct {
	ID       int             `json:"id"`
	Name     TaskType        `json:"name"`
	Created  time.Time       `json:"created"`
	Started  *time.Time      `json:"start_at"`
	Finished *time.Time      `json:"finish_at"`
	Updated  *time.Time      `json:"-"`
	Status   string          `json:"status"`
	User     *userV1.User    `json:"-"`
	Payload  json.RawMessage `json:"payload"`
	Result   json.RawMessage `json:"result"`
}
