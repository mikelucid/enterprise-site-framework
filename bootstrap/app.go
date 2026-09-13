package bootstrap

import (
	"fmt"
	"reflect"
	"sync"
)

type Lifetime int

const (
	Transient Lifetime = iota
	Singleton
)

type ServiceFactory func(c *ServiceContainer) (any, error)

type registration struct {
	lifetime Lifetime
	factory  ServiceFactory
	instance any
}

type ServiceContainer struct {
	mu       sync.RWMutex
	services map[string]*registration
}

var globalContainer = NewContainer()

func NewContainer() *ServiceContainer {
	return &ServiceContainer{services: map[string]*registration{}}
}

func GlobalContainer() *ServiceContainer { return globalContainer }

func SetGlobalContainer(c *ServiceContainer) {
	if c == nil {
		return
	}
	globalContainer = c
}

func (c *ServiceContainer) Register(name string, lifetime Lifetime, factory ServiceFactory) error {
	if name == "" || factory == nil {
		return fmt.Errorf("invalid service registration: name and factory are required")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.services[name] = &registration{lifetime: lifetime, factory: factory}
	return nil
}

func (c *ServiceContainer) RegisterInstance(name string, instance any) error {
	if name == "" || instance == nil {
		return fmt.Errorf("invalid service instance registration")
	}
	return c.Register(name, Singleton, func(*ServiceContainer) (any, error) { return instance, nil })
}

func (c *ServiceContainer) Resolve(name string) (any, error) {
	c.mu.RLock()
	reg, ok := c.services[name]
	c.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("service %q is not registered", name)
	}
	if reg.lifetime == Singleton {
		c.mu.Lock()
		defer c.mu.Unlock()
		if reg.instance == nil {
			instance, err := reg.factory(c)
			if err != nil {
				return nil, fmt.Errorf("failed creating singleton %q: %w", name, err)
			}
			reg.instance = instance
		}
		return reg.instance, nil
	}
	instance, err := reg.factory(c)
	if err != nil {
		return nil, fmt.Errorf("failed creating transient %q: %w", name, err)
	}
	return instance, nil
}

func ResolveAs[T any](c *ServiceContainer, name string) (T, error) {
	var zero T
	service, err := c.Resolve(name)
	if err != nil {
		return zero, err
	}
	casted, ok := service.(T)
	if ok {
		return casted, nil
	}
	t := reflect.TypeOf((*T)(nil)).Elem()
	return zero, fmt.Errorf("service %q is %T, expected %s", name, service, t.String())
}
