package domain

import "context"

type Repository interface {
	Create(context.Context, Site) (Site, error)
	Get(context.Context, string) (Site, error)
	List(context.Context) ([]Site, error)
	Update(context.Context, Site) (Site, error)
	Delete(context.Context, string) error
}
