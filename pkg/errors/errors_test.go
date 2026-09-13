package errors

import "testing"

func TestValidationError(t *testing.T) {
	err := ValidationError("bad")
	if err.StatusCode != 400 {
		t.Fatalf("unexpected status: %d", err.StatusCode)
	}
}
