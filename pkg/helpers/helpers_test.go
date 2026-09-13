package helpers

import "testing"

func TestNewUUID(t *testing.T) {
	if NewUUID() == "" {
		t.Fatal("expected uuid")
	}
}
