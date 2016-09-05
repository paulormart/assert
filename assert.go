package assert

import (
	"reflect"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
}

// Equal asserts two objects are deeply equal
func Equal(t TestingT, expected, actual interface{}) bool {

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal: %v (expected) != %v", expected, actual)
		return false
	}
	return true
}


