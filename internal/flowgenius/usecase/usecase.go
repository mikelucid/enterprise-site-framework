// Package usecase orchestrates composition, release and royalty flows.
package usecase

import (
	"context"
	"fmt"

	"github.com/mikelucid/enterprise-site-framework/internal/flowgenius/domain"
)

// Composer generates music (AI model adapter).
type Composer interface {
	Compose(ctx context.Context, req domain.CompositionRequest) (domain.Track, error)
}

// Beatport is the marketplace adapter (Beatport partner/label API).
type Beatport interface {
	SubmitRelease(ctx context.Context, label domain.Label, rel domain.Release) (beatportID string, err error)
}

// Repository persists labels, tracks and releases.
type Repository interface {
	GetLabel(ctx context.Context, id string) (domain.Label, error)
	SaveTrack(ctx context.Context, t domain.Track) error
	GetTrack(ctx context.Context, id string) (domain.Track, error)
	SaveRelease(ctx context.Context, r domain.Release) error
	GetRelease(ctx context.Context, id string) (domain.Release, error)
}

type Service struct {
	Composer Composer
	Beatport Beatport
	Repo     Repository
}

// GenerateTrack composes a track and stores it with the given splits.
// If splits is empty the creator receives 100%.
func (s *Service) GenerateTrack(ctx context.Context, req domain.CompositionRequest, splits []domain.Split) (domain.Track, error) {
	if len(splits) == 0 {
		splits = []domain.Split{{PayeeID: req.CreatorID, BPS: 10000}}
	}
	if err := domain.ValidateSplits(splits); err != nil {
		return domain.Track{}, err
	}
	t, err := s.Composer.Compose(ctx, req)
	if err != nil {
		return domain.Track{}, fmt.Errorf("compose: %w", err)
	}
	t.Splits = splits
	if err := s.Repo.SaveTrack(ctx, t); err != nil {
		return domain.Track{}, err
	}
	return t, nil
}

// CreateRelease bundles tracks into a draft release for a label.
func (s *Service) CreateRelease(ctx context.Context, id, labelID, title string, trackIDs []string) (domain.Release, error) {
	if labelID == "" {
		return domain.Release{}, domain.ErrNoLabel
	}
	if len(trackIDs) == 0 {
		return domain.Release{}, domain.ErrNoTracks
	}
	if _, err := s.Repo.GetLabel(ctx, labelID); err != nil {
		return domain.Release{}, err
	}
	r := domain.Release{ID: id, LabelID: labelID, Title: title, Status: domain.ReleaseDraft}
	for _, tid := range trackIDs {
		t, err := s.Repo.GetTrack(ctx, tid)
		if err != nil {
			return domain.Release{}, err
		}
		r.Tracks = append(r.Tracks, t)
	}
	return r, s.Repo.SaveRelease(ctx, r)
}

// SubmitToBeatport sends a draft release to the marketplace.
func (s *Service) SubmitToBeatport(ctx context.Context, releaseID string) (domain.Release, error) {
	r, err := s.Repo.GetRelease(ctx, releaseID)
	if err != nil {
		return domain.Release{}, err
	}
	if r.Status != domain.ReleaseDraft {
		return domain.Release{}, domain.ErrNotDraft
	}
	l, err := s.Repo.GetLabel(ctx, r.LabelID)
	if err != nil {
		return domain.Release{}, err
	}
	bpID, err := s.Beatport.SubmitRelease(ctx, l, r)
	if err != nil {
		r.Status = domain.ReleaseRejected
		_ = s.Repo.SaveRelease(ctx, r)
		return r, fmt.Errorf("beatport submit: %w", err)
	}
	r.Status, r.BeatportID = domain.ReleaseSubmitted, bpID
	return r, s.Repo.SaveRelease(ctx, r)
}

// SplitSale allocates a sale amount (minor units) across payees of every
// track in the release, dividing evenly per track.
func (s *Service) SplitSale(ctx context.Context, releaseID string, amount int64) (map[string]int64, error) {
	r, err := s.Repo.GetRelease(ctx, releaseID)
	if err != nil {
		return nil, err
	}
	if len(r.Tracks) == 0 {
		return nil, domain.ErrNoTracks
	}
	n := int64(len(r.Tracks))
	out := map[string]int64{}
	for i, t := range r.Tracks {
		share := amount / n
		if int64(i) < amount%n {
			share++
		}
		a, err := domain.Allocate(share, t.Splits)
		if err != nil {
			return nil, err
		}
		for k, v := range a {
			out[k] += v
		}
	}
	return out, nil
}
