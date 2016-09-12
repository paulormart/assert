package assert_test

import (
	"testing"
	"assert"
)

type User struct {
	name 	string
}

func TestAll(t *testing.T){
	// Table tests
	var tests = []struct{
		expected interface{}
		actual interface{}
		result bool
	}{{
		expected: "Hello World",
		actual: "Hello World",
		result: true,
	},
	{
		expected: nil,
		actual: nil,
		result: true,
	},
	{
		expected: func()*User{
			return &User{}
		}(),
		actual: &User{},
		result: true,
	}}
	//TODO add more tests

	for _, test := range tests {
		if res := assert.Equal(t, test.expected, test.actual); res != test.result {
			t.Errorf("assert.Equal(%v, %v) = %v", test.expected, test.actual, test.result)
		}

	}
}
