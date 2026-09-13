package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	JWT struct {
		Secret     string `mapstructure:"secret"`
		Algorithm  string `mapstructure:"algorithm"`
		Expiration int64  `mapstructure:"expiration"`
	} `mapstructure:"jwt"`
	RBAC struct {
		DefaultRole string   `mapstructure:"default_role"`
		Roles       []string `mapstructure:"roles"`
	} `mapstructure:"rbac"`
}

type Claims struct {
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

type Manager struct {
	cfg        Config
	hsSecret   []byte
	rsPrivate  *rsa.PrivateKey
	rsPublic   *rsa.PublicKey
	revocation map[string]struct{}
}

func NewManager(cfg Config) (*Manager, error) {
	m := &Manager{cfg: cfg, hsSecret: []byte(cfg.JWT.Secret), revocation: map[string]struct{}{}}
	if cfg.JWT.Algorithm == "RS256" {
		key, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, err
		}
		m.rsPrivate = key
		m.rsPublic = &key.PublicKey
	}
	if cfg.JWT.Expiration <= 0 {
		m.cfg.JWT.Expiration = 86400
	}
	if m.cfg.JWT.Algorithm == "" {
		m.cfg.JWT.Algorithm = "HS256"
	}
	return m, nil
}

func (m *Manager) GenerateToken(subject string, roles []string) (string, error) {
	now := time.Now()
	claims := Claims{Roles: roles, RegisteredClaims: jwt.RegisteredClaims{Subject: subject, IssuedAt: jwt.NewNumericDate(now), ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(m.cfg.JWT.Expiration) * time.Second))}}
	var method jwt.SigningMethod = jwt.SigningMethodHS256
	key := any(m.hsSecret)
	if m.cfg.JWT.Algorithm == "RS256" {
		method = jwt.SigningMethodRS256
		key = m.rsPrivate
	}
	return jwt.NewWithClaims(method, claims).SignedString(key)
}

func (m *Manager) ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if m.cfg.JWT.Algorithm == "RS256" {
			return m.rsPublic, nil
		}
		return m.hsSecret, nil
	})
	if err != nil || !parsed.Valid {
		return nil, errors.New("invalid token")
	}
	if _, revoked := m.revocation[claims.ID]; revoked {
		return nil, errors.New("token revoked")
	}
	return claims, nil
}

func (m *Manager) RefreshToken(token string) (string, error) {
	claims, err := m.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return m.GenerateToken(claims.Subject, claims.Roles)
}

func (m *Manager) RevokeToken(jti string) { m.revocation[jti] = struct{}{} }
