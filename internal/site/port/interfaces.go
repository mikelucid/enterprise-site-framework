package port

import (
	"context"

	"github.com/mikelucid/enterprise-site-framework/internal/site/domain"
)

type Service interface {
	Create(context.Context, domain.Site) (domain.Site, error)
	Get(context.Context, string) (domain.Site, error)
}
