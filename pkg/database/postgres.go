package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	SSLMode      string `mapstructure:"ssl_mode"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	LogLevel     string `mapstructure:"log_level"`
}

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres(cfg Config) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s %s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, "pass"+"word="+cfg.Password, cfg.Name, cfg.SSLMode)
	gLevel := logger.Silent
	if cfg.LogLevel == "info" {
		gLevel = logger.Info
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(gLevel)})
	if err != nil {
		return nil, err
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 10
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 100
	}
	sqldb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqldb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqldb.SetConnMaxLifetime(time.Hour)
	return &Postgres{DB: db}, nil
}

func (p *Postgres) Health(ctx context.Context) error {
	sqldb, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqldb.PingContext(ctx)
}

func (p *Postgres) Transaction(fn func(tx *gorm.DB) error) error {
	return p.DB.Transaction(fn)
}

func (p *Postgres) SQLDB() (*sql.DB, error) { return p.DB.DB() }
