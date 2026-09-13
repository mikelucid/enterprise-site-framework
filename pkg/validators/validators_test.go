package validators

import "testing"

func TestValidators(t *testing.T) {
	if !Email("hello@example.com") {
		t.Fatal("expected valid email")
	}
	if !URL("https://example.com") {
		t.Fatal("expected valid url")
	}
	if !UUID("550e8400-e29b-41d4-a716-446655440000") {
		t.Fatal("expected valid uuid")
	}
	if !Phone("+14155552671") {
		t.Fatal("expected valid phone")
	}
}
