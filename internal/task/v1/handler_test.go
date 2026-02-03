package task

import (
	"testing"
)

type TaskTestCase struct {
	Name string
}

func TestTaskHandle(t *testing.T) {
	testCases := []*TaskTestCase{
		&TaskTestCase{},
		&TaskTestCase{},
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
