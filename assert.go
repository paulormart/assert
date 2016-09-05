package assert

import (
	"reflect"
	"testing"
)

// Equal asserts two objects are deeply equal
func Equal(t *testing.T, expected, actual interface{}) bool {

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal: %v (expected) != %v", expected, actual)
		return false
	}
	return true
}


