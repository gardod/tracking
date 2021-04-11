package service

import (
	"tratnik.net/client/internal/model"
)

func ExampleMessage_process_filterNone() {
	testCases := []model.Message{
		{AccountID: 1, Data: []byte("foo")},
		{AccountID: 2, Data: []byte("bar")},
		{AccountID: 0, Data: []byte("baz")},
	}

	messageSrvc := &Message{}
	process := messageSrvc.process(model.Filter{})

	for _, testCase := range testCases {
		process(testCase)
	}

	// Output:
	// AccountID:   1    Timestamp: 0001-01-01 00:00:00.000 +0000    Data: foo
	// AccountID:   2    Timestamp: 0001-01-01 00:00:00.000 +0000    Data: bar
	// AccountID:   0    Timestamp: 0001-01-01 00:00:00.000 +0000    Data: baz
}

func ExampleMessage_process_filterAccountID() {
	testCases := []model.Message{
		{AccountID: 1, Data: []byte("foo")},
		{AccountID: 2, Data: []byte("bar")},
		{AccountID: 0, Data: []byte("baz")},
		{AccountID: 1, Data: []byte("qux")},
	}

	messageSrvc := &Message{}
	process := messageSrvc.process(model.Filter{AccountID: 1})

	for _, testCase := range testCases {
		process(testCase)
	}

	// Output:
	// AccountID:   1    Timestamp: 0001-01-01 00:00:00.000 +0000    Data: foo
	// AccountID:   1    Timestamp: 0001-01-01 00:00:00.000 +0000    Data: qux
}
