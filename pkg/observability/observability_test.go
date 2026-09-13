package observability

import "testing"

func TestMeter(t *testing.T) {
	if Meter("x") == nil {
		t.Fatal("expected meter")
	}
}
