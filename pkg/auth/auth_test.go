package auth

import "testing"

func TestTokenLifecycle(t *testing.T) {
	cfg := Config{}
	cfg.JWT.Secret = "secret"
	cfg.JWT.Algorithm = "HS256"
	mgr, err := NewManager(cfg)
	if err != nil {
		t.Fatal(err)
	}
	token, err := mgr.GenerateToken("user-1", []string{"admin"})
	if err != nil {
		t.Fatal(err)
	}
	claims, err := mgr.ValidateToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if claims.Subject != "user-1" {
		t.Fatalf("unexpected subject: %s", claims.Subject)
	}
}
