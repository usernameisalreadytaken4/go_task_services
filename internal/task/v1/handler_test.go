package task

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/usernameisalreadytaken4/go_task_services/internal"
	userV1 "github.com/usernameisalreadytaken4/go_task_services/internal/user/v1"
)

type TaskTestCase struct {
	Name               string
	Method             string
	URL                string
	Token              string
	ResponseHTTPStatus int
	Payload            json.RawMessage
	Expected           interface{}
}

type TestTask struct {
	ID       int             `json:"id"`
	Name     TaskType        `json:"name"`
	Created  time.Time       `json:"created"`
	Started  time.Time       `json:"start_at"`
	Finished time.Time       `json:"finish_at"`
	Status   string          `json:"status"`
	Payload  json.RawMessage `json:"payload"`
	Result   json.RawMessage `json:"result"`
}

func TestTaskHandle(t *testing.T) {

	// проверяю как работают моки
	// в целом, тут это стрельба из пушек по воробьям, много нагородил, чтобы показать что возвращается Created
	// но в более сложных сервисах с более сложной логикой, или в тестировании самих сервисов это будет эффективнее

	shortTaskPayload, _ := json.Marshal(map[string]string{
		"type": "short_task",
		"text": "Nobody cares",
	})

	// longTaskPayload, _ := json.Marshal(map[string]string{
	// 	"type": "long_task",
	// 	"text": "Everybody cares",
	// })

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedService := NewMockService(ctrl)
	handler := &Handler{
		service: mockedService,
	}

	var body io.Reader = bytes.NewReader(shortTaskPayload)

	testUser := &userV1.User{
		ID: 1,
	}

	req := httptest.NewRequest("POST", "/api/v1/tasks", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer testtoken")

	req = req.WithContext(internal.ContextWithUser(req.Context(), testUser))

	w := httptest.NewRecorder()

	mockedService.EXPECT().
		Create(testUser, gomock.Any()).
		DoAndReturn(func(u *userV1.User, task *Task) (*Task, error) {
			if task.Name != "short_task" {
				t.Errorf("wrong name\nEXPECTED: %v\nGET: %v\n", "short_task", task.Name)
			}
			return &Task{ID: 1}, nil
		})

	handler.Handle(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("wrong status\nEXPECTED: %v\nGET: %v\n", http.StatusCreated, w.Code)
	}
}
