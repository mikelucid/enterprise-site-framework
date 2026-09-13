package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Loader struct {
	viper *viper.Viper
}

type Resolver interface {
	Resolve(name string) (any, error)
}

func New(paths ...string) *Loader {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	for _, p := range paths {
		v.AddConfigPath(p)
	}
	return &Loader{viper: v}
}

func (l *Loader) ReadConfig(name string) error {
	l.viper.SetConfigName(name)
	return l.viper.ReadInConfig()
}

func (l *Loader) Unmarshal(key string, out any) error {
	if !l.viper.IsSet(key) {
		return fmt.Errorf("config key %q not set", key)
	}
	return l.viper.UnmarshalKey(key, out)
}

func (l *Loader) Watch(onChange func(fsnotify.Event)) {
	l.viper.OnConfigChange(onChange)
	l.viper.WatchConfig()
}

func Get[T any](resolver Resolver, key string) (*T, error) {
	svc, err := resolver.Resolve(key)
	if err != nil {
		return nil, err
	}
	cfg, ok := svc.(*T)
	if !ok {
		v, ok := svc.(T)
		if !ok {
			return nil, fmt.Errorf("service %q is %T, expected %T", key, svc, *new(T))
		}
		return &v, nil
	}
	return cfg, nil
}
