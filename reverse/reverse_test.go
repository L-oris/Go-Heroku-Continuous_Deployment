package reverse

import (
	"reflect"
	"testing"
)

func TestReverseToReturnString(t *testing.T) {
	result := Reverse("hello")

	if reflect.TypeOf(result).Kind() != reflect.String {
		t.Fatalf("Expected string, got %T", result)
	}
}

func TestReverseToReturnReversedInputString(t *testing.T) {
	result := Reverse("hello")
	expectedResult := "ollh"

	if result != expectedResult {
		t.Fatalf("Expected %s, got %s", result, expectedResult)
	}
}
