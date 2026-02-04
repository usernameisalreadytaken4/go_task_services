package task

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
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

	req := httptest.NewRequest("POST", "/api/v1/tasks", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer testtoken")

	w := httptest.NewRecorder()

	handler.Handle(w, req)

	resp := w.Result()

	log.Println(resp)

	// mockedService.EXPECT().GetByUserID(1).Return(result, nil)

	// log.Panicln(ts)

	// testCases := []*TaskTestCase{
	// 	{
	// 		Name:               "Create Short Task",
	// 		Method:             "POST",
	// 		URL:                "/api/v1/tasks",
	// 		ResponseHTTPStatus: 201,
	// 		Payload:            shortTaskPayload,
	// 		Expected: func() interface{} {
	// 			return &struct {
	// 				Task TestTask
	// 			}{
	// 				Task: TestTask{
	// 					Name:   "short_task",
	// 					Status: "new",
	// 				},
	// 			}
	// 		},
	// 	},
	// 	{
	// 		Name:               "Create Long Task",
	// 		Method:             "POST",
	// 		URL:                "/api/v1/tasks",
	// 		ResponseHTTPStatus: 201,
	// 		Payload:            longTaskPayload,
	// 		Expected: func() interface{} {
	// 			return &struct {
	// 				Task TestTask
	// 			}{
	// 				Task: TestTask{
	// 					Name:   "long_task",
	// 					Status: "new",
	// 				},
	// 			}
	// 		},
	// 	},
	// }

	// for _, testCase := range testCases {
	// 	ok := t.Run(testCase.Name, func(t *testing.T) {
	// 		// make some things
	// 	})
	// 	if !ok {
	// 		break
	// 	}
	// }
}
