package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Config struct {
	JWT struct {
		Secret     string `mapstructure:"secret"`
		Algorithm  string `mapstructure:"algorithm"`
		Expiration int64  `mapstructure:"expiration"`
		PrivateKey string `mapstructure:"private_key"`
		PublicKey  string `mapstructure:"public_key"`
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
	revMu      sync.RWMutex
}

func NewManager(cfg Config) (*Manager, error) {
	m := &Manager{cfg: cfg, hsSecret: []byte(cfg.JWT.Secret), revocation: map[string]struct{}{}}
	if m.cfg.JWT.Algorithm == "" {
		m.cfg.JWT.Algorithm = "HS256"
	}
	if m.cfg.JWT.Algorithm == "HS256" && len(m.hsSecret) == 0 {
		return nil, errors.New("jwt secret is required for HS256")
	}
	if m.cfg.JWT.Algorithm == "RS256" {
		if m.cfg.JWT.PrivateKey != "" && m.cfg.JWT.PublicKey != "" {
			priv, pub, err := parseRSAKeys(m.cfg.JWT.PrivateKey, m.cfg.JWT.PublicKey)
			if err != nil {
				return nil, err
			}
			m.rsPrivate, m.rsPublic = priv, pub
		} else {
			key, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				return nil, err
			}
			m.rsPrivate = key
			m.rsPublic = &key.PublicKey
		}
	}
	if cfg.JWT.Expiration <= 0 {
		m.cfg.JWT.Expiration = 86400
	}
	return m, nil
}

func (m *Manager) GenerateToken(subject string, roles []string) (string, error) {
	now := time.Now()
	claims := Claims{
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(m.cfg.JWT.Expiration) * time.Second)),
		},
	}
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
		if t.Method.Alg() != m.cfg.JWT.Algorithm {
			return nil, errors.New("unexpected signing algorithm")
		}
		if m.cfg.JWT.Algorithm == "RS256" {
			return m.rsPublic, nil
		}
		return m.hsSecret, nil
	})
	if err != nil || !parsed.Valid {
		return nil, errors.New("invalid token")
	}
	m.revMu.RLock()
	_, revoked := m.revocation[claims.ID]
	m.revMu.RUnlock()
	if revoked {
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

func (m *Manager) RevokeToken(jti string) {
	m.revMu.Lock()
	m.revocation[jti] = struct{}{}
	m.revMu.Unlock()
}

func parseRSAKeys(privatePEM, publicPEM string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privBlock, _ := pem.Decode([]byte(privatePEM))
	if privBlock == nil {
		return nil, nil, fmt.Errorf("invalid private key PEM")
	}
	privAny, err := x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}
	priv, ok := privAny.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("private key is not RSA")
	}
	pubBlock, _ := pem.Decode([]byte(publicPEM))
	if pubBlock == nil {
		return nil, nil, fmt.Errorf("invalid public key PEM")
	}
	pubAny, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}
	pub, ok := pubAny.(*rsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("public key is not RSA")
	}
	return priv, pub, nil
}
