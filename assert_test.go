package assert_test

import (
	"testing"

	"github.com/paulormart/assert"
)

type FakeStruct struct {
	SomeSlice       []string
	SomeStruct      FakeInnerStruct
	SomeSliceStruct []FakeInnerStruct
}
type FakeInnerStruct struct {
	ID      string
	SomeMap map[int]string
}

func TestAll(t *testing.T) {
	// Table tests
	var tests = []struct {
		expected interface{}
		actual   interface{}
		result   bool
	}{
		{
			expected: "Hello World",
			actual:   "Hello World",
			result:   true,
		},
		{
			expected: nil,
			actual:   nil,
			result:   true,
		},
		{
			expected: func() *FakeStruct {
				return &FakeStruct{}
			}(),
			actual: &FakeStruct{},
			result: true,
		},
		{
			expected: map[string]string{"key1": "foo", "key2": "bar"},
			actual:   map[string]string{"key1": "foo", "key2": "bar"},
			result:   true,
		},
		{
			expected: FakeStruct{
				SomeSlice:       []string{"foo", "bar"},
				SomeStruct:      FakeInnerStruct{ID: "abc", SomeMap: map[int]string{1: "foo", 2: "bar"}},
				SomeSliceStruct: []FakeInnerStruct{{ID: "abc", SomeMap: map[int]string{1: "foo", 2: "bar"}}}},
			actual: FakeStruct{
				SomeSlice:       []string{"foo", "bar"},
				SomeStruct:      FakeInnerStruct{"abc", map[int]string{1: "foo", 2: "bar"}},
				SomeSliceStruct: []FakeInnerStruct{{"abc", map[int]string{1: "foo", 2: "bar"}}}},
			result: true,
		},
	}
	//TODO add more tests

	for _, test := range tests {
		if res := assert.Equal(t, test.expected, test.actual); res != test.result {
			t.Errorf("assert.Equal(%v, %v) = %v", test.expected, test.actual, test.result)
		}

	}
}
