package adapter

import (
	"context"

	"github.com/mikelucid/enterprise-site-framework/internal/ai/domain"
)

type OpenAIModel struct{}

func (OpenAIModel) Recommend(context.Context, string, int) ([]domain.Recommendation, error) {
	return []domain.Recommendation{}, nil
}
