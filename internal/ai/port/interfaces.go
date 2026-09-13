package port

import (
	"context"

	"github.com/mikelucid/enterprise-site-framework/internal/ai/domain"
)

type Service interface {
	GetRecommendations(context.Context, string, int) ([]domain.Recommendation, error)
}
