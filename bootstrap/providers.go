package bootstrap

import (
	"fmt"

	"github.com/mikelucid/enterprise-site-framework/pkg/ai"
	"github.com/mikelucid/enterprise-site-framework/pkg/auth"
	"github.com/mikelucid/enterprise-site-framework/pkg/cache"
	"github.com/mikelucid/enterprise-site-framework/pkg/config"
	"github.com/mikelucid/enterprise-site-framework/pkg/database"
	"github.com/mikelucid/enterprise-site-framework/pkg/events"
	"github.com/mikelucid/enterprise-site-framework/pkg/logger"
	"github.com/mikelucid/enterprise-site-framework/pkg/messaging"
)

type Provider interface {
	Register(*ServiceContainer) error
	Name() string
}

type DatabaseProvider struct{}
type LoggerProvider struct{}
type AIProvider struct{}
type QueueProvider struct{}
type EventProvider struct{}
type CacheProvider struct{}
type AuthProvider struct{}

func (DatabaseProvider) Name() string { return "database" }
func (LoggerProvider) Name() string   { return "logger" }
func (AIProvider) Name() string       { return "ai" }
func (QueueProvider) Name() string    { return "queue" }
func (EventProvider) Name() string    { return "events" }
func (CacheProvider) Name() string    { return "cache" }
func (AuthProvider) Name() string     { return "auth" }

func (DatabaseProvider) Register(c *ServiceContainer) error {
	return c.Register("database", Singleton, func(c *ServiceContainer) (any, error) {
		cfg, err := config.Get[database.Config](c, "config.database")
		if err != nil {
			return nil, err
		}
		return database.NewPostgres(*cfg)
	})
}

func (LoggerProvider) Register(c *ServiceContainer) error {
	return c.Register("logger", Singleton, func(c *ServiceContainer) (any, error) {
		cfg, err := config.Get[logger.Config](c, "config.logger")
		if err != nil {
			cfg = &logger.Config{Level: "info"}
		}
		return logger.New(*cfg)
	})
}

func (AIProvider) Register(c *ServiceContainer) error {
	return c.Register("ai", Singleton, func(c *ServiceContainer) (any, error) { return ai.NewEngine(), nil })
}

func (QueueProvider) Register(c *ServiceContainer) error {
	return c.Register("queue", Singleton, func(c *ServiceContainer) (any, error) {
		cfg, err := config.Get[messaging.Config](c, "config.queue")
		if err != nil {
			return nil, err
		}
		return messaging.NewClient(*cfg)
	})
}

func (EventProvider) Register(c *ServiceContainer) error {
	return c.Register("events", Singleton, func(*ServiceContainer) (any, error) { return events.NewDispatcher(), nil })
}

func (CacheProvider) Register(c *ServiceContainer) error {
	return c.Register("cache", Singleton, func(*ServiceContainer) (any, error) { return cache.NewInMemory(), nil })
}

func (AuthProvider) Register(c *ServiceContainer) error {
	return c.Register("auth", Singleton, func(c *ServiceContainer) (any, error) {
		cfg, err := config.Get[auth.Config](c, "config.auth")
		if err != nil {
			return nil, err
		}
		return auth.NewManager(*cfg)
	})
}

func RegisterProviders(c *ServiceContainer, providers ...Provider) error {
	for _, p := range providers {
		if err := p.Register(c); err != nil {
			return fmt.Errorf("provider %s failed: %w", p.Name(), err)
		}
	}
	return nil
}

func RegisterDefaultProviders(c *ServiceContainer) error {
	return RegisterProviders(c,
		LoggerProvider{},
		DatabaseProvider{},
		AIProvider{},
		QueueProvider{},
		EventProvider{},
		CacheProvider{},
		AuthProvider{},
	)
}
