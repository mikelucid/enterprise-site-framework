package bootstrap

import "testing"

func TestContainerSingleton(t *testing.T) {
	c := NewContainer()
	if err := c.Register("x", Singleton, func(*ServiceContainer) (any, error) { return &struct{}{}, nil }); err != nil {
		t.Fatal(err)
	}
	a, _ := c.Resolve("x")
	b, _ := c.Resolve("x")
	if a != b {
		t.Fatal("expected singleton instance")
	}
}
