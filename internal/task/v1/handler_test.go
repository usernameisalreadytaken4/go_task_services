package task

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

	shortTaskPayload, _ := json.Marshal(map[string]string{
		"type": "short_task",
		"text": "Nobody cares",
	})

	longTaskPayload, _ := json.Marshal(map[string]string{
		"type": "long_task",
		"text": "Everybody cares",
	})

	dsn := "postgres://postgres:my_password@localhost:5432/task_service"
	// temporary. Will be changed to mocks
	config, _ := pgxpool.ParseConfig(dsn)
	pool, _ := pgxpool.NewWithConfig(context.Background(), config)
	err := pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	taskRepo := NewRepository(pool)
	taskService := NewService(taskRepo)
	taskHandler := NewHandler(taskService)

	TaskRouter(mux, taskHandler)
	ts := httptest.NewServer(mux)

	log.Panicln(ts)

	testCases := []*TaskTestCase{
		&TaskTestCase{
			Name:               "Create Short Task",
			Method:             "POST",
			URL:                "/api/v1/tasks",
			ResponseHTTPStatus: 201,
			Payload:            shortTaskPayload,
			Expected: func() interface{} {
				return &struct {
					Task TestTask
				}{
					Task: TestTask{
						Name:   "short_task",
						Status: "new",
					},
				}
			},
		},
		&TaskTestCase{
			Name:               "Create Long Task",
			Method:             "POST",
			URL:                "/api/v1/tasks",
			ResponseHTTPStatus: 201,
			Payload:            longTaskPayload,
			Expected: func() interface{} {
				return &struct {
					Task TestTask
				}{
					Task: TestTask{
						Name:   "long_task",
						Status: "new",
					},
				}
			},
		},
	}

	for _, testCase := range testCases {
		ok := t.Run(testCase.Name, func(t *testing.T) {
			// make some things
		})
		if !ok {
			break
		}
	}
}
