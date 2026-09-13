package config

import "testing"

type fakeResolver map[string]any

func (f fakeResolver) Resolve(name string) (any, error) { return f[name], nil }

func TestGet(t *testing.T) {
	type sample struct{ Name string }
	cfg, err := Get[sample](fakeResolver{"x": sample{Name: "ok"}}, "x")
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Name != "ok" {
		t.Fatalf("unexpected value: %s", cfg.Name)
	}
}
