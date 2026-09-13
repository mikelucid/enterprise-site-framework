package adapter

import (
	"context"

	"github.com/mikelucid/enterprise-site-framework/internal/site/domain"
)

type PostgresRepository struct{}

func (PostgresRepository) Create(context.Context, domain.Site) (domain.Site, error) { return domain.Site{}, nil }
func (PostgresRepository) Get(context.Context, string) (domain.Site, error)          { return domain.Site{}, nil }
func (PostgresRepository) List(context.Context) ([]domain.Site, error)                { return []domain.Site{}, nil }
func (PostgresRepository) Update(context.Context, domain.Site) (domain.Site, error)   { return domain.Site{}, nil }
func (PostgresRepository) Delete(context.Context, string) error                        { return nil }
