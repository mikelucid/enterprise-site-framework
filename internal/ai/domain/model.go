package domain

import "context"

type Model interface {
	Recommend(context.Context, string, int) ([]Recommendation, error)
}
