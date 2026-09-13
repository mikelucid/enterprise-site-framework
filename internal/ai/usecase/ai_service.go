package usecase

import (
	"context"

	"github.com/mikelucid/enterprise-site-framework/internal/ai/domain"
)

type Service struct{ model domain.Model }

func New(model domain.Model) *Service { return &Service{model: model} }

func (s *Service) GetRecommendations(ctx context.Context, siteID string, limit int) ([]domain.Recommendation, error) {
	return s.model.Recommend(ctx, siteID, limit)
}
