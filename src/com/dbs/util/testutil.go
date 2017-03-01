package util

import (
	"fmt"
	"testing"
)


func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func AssertTrue(t *testing.T, condition bool, message string) {
	if condition {
		return
	} else {
		if len(message) == 0 {
			message = fmt.Sprintf("condition mismatch")
		}
		t.Fatal(message)
	}
}

