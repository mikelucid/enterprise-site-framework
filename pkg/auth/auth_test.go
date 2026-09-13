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
	refreshed, err := mgr.RefreshToken(token)
	if err != nil {
		t.Fatal(err)
	}
	refClaims, err := mgr.ValidateToken(refreshed)
	if err != nil {
		t.Fatal(err)
	}
	if refClaims.Subject != claims.Subject {
		t.Fatalf("subject mismatch after refresh: %s != %s", refClaims.Subject, claims.Subject)
	}
	mgr.RevokeToken(refClaims.ID)
	if _, err := mgr.ValidateToken(refreshed); err == nil {
		t.Fatal("expected revoked token to fail validation")
	}
}

func TestTokenLifecycleRS256(t *testing.T) {
	cfg := Config{}
	cfg.JWT.Algorithm = "RS256"
	mgr, err := NewManager(cfg)
	if err != nil {
		t.Fatal(err)
	}
	token, err := mgr.GenerateToken("user-rs", []string{"manager"})
	if err != nil {
		t.Fatal(err)
	}
	claims, err := mgr.ValidateToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if claims.Subject != "user-rs" {
		t.Fatalf("unexpected subject: %s", claims.Subject)
	}
}

func TestAlgorithmMismatchRejected(t *testing.T) {
	hsCfg := Config{}
	hsCfg.JWT.Secret = "secret"
	hsCfg.JWT.Algorithm = "HS256"
	hsMgr, err := NewManager(hsCfg)
	if err != nil {
		t.Fatal(err)
	}
	token, err := hsMgr.GenerateToken("user-mismatch", []string{"user"})
	if err != nil {
		t.Fatal(err)
	}

	rsCfg := Config{}
	rsCfg.JWT.Algorithm = "RS256"
	rsMgr, err := NewManager(rsCfg)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := rsMgr.ValidateToken(token); err == nil {
		t.Fatal("expected algorithm mismatch to fail")
	}
}
