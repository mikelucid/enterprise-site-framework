package database

import "testing"

func TestConfigDefaults(t *testing.T) {
	cfg := Config{}
	if cfg.MaxIdleConns != 0 || cfg.MaxOpenConns != 0 {
		t.Fatal("expected zero defaults before init")
	}
}
