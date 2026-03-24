package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("ENV", "")
	t.Setenv("PORT", "")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("JWT_SECRET", "")

	cfg := Load()

	if cfg.Env != "development" {
		t.Fatalf("expected default env, got %q", cfg.Env)
	}
	if cfg.Port != "8080" {
		t.Fatalf("expected default port, got %q", cfg.Port)
	}
	if cfg.DBConn != "postgres://localhost/stellarbill?sslmode=disable" {
		t.Fatalf("expected default db conn, got %q", cfg.DBConn)
	}
	if cfg.JWTSecret != "change-me-in-production" {
		t.Fatalf("expected default jwt secret, got %q", cfg.JWTSecret)
	}
}

func TestLoadEnvOverrides(t *testing.T) {
	t.Setenv("ENV", "production")
	t.Setenv("PORT", "9090")
	t.Setenv("DATABASE_URL", "postgres://example/db")
	t.Setenv("JWT_SECRET", "super-secret")

	cfg := Load()

	if cfg.Env != "production" || cfg.Port != "9090" || cfg.DBConn != "postgres://example/db" || cfg.JWTSecret != "super-secret" {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}
