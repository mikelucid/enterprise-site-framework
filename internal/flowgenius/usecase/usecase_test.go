package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/mikelucid/enterprise-site-framework/internal/flowgenius/domain"
)

type fakeComposer struct{}

func (fakeComposer) Compose(_ context.Context, r domain.CompositionRequest) (domain.Track, error) {
	return domain.Track{ID: "t1", Title: "AI " + r.Genre, Genre: r.Genre, BPM: r.BPM}, nil
}

type fakeBP struct{ err error }

func (f fakeBP) SubmitRelease(context.Context, domain.Label, domain.Release) (string, error) {
	return "bp-1", f.err
}

type memRepo struct {
	tracks   map[string]domain.Track
	releases map[string]domain.Release
}

func (m *memRepo) GetLabel(_ context.Context, id string) (domain.Label, error) {
	return domain.Label{ID: id, Name: "L"}, nil
}
func (m *memRepo) SaveTrack(_ context.Context, t domain.Track) error { m.tracks[t.ID] = t; return nil }
func (m *memRepo) GetTrack(_ context.Context, id string) (domain.Track, error) {
	t, ok := m.tracks[id]
	if !ok {
		return t, errors.New("not found")
	}
	return t, nil
}
func (m *memRepo) SaveRelease(_ context.Context, r domain.Release) error {
	m.releases[r.ID] = r
	return nil
}
func (m *memRepo) GetRelease(_ context.Context, id string) (domain.Release, error) {
	r, ok := m.releases[id]
	if !ok {
		return r, errors.New("not found")
	}
	return r, nil
}

func newSvc(bp fakeBP) *Service {
	return &Service{fakeComposer{}, bp, &memRepo{map[string]domain.Track{}, map[string]domain.Release{}}}
}

func TestFlow(t *testing.T) {
	ctx := context.Background()
	s := newSvc(fakeBP{})
	_, err := s.GenerateTrack(ctx, domain.CompositionRequest{Genre: "Techno", CreatorID: "a"},
		[]domain.Split{{PayeeID: "a", BPS: 7000}, {PayeeID: "b", BPS: 3000}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.CreateRelease(ctx, "r1", "label", "EP", []string{"t1"}); err != nil {
		t.Fatal(err)
	}
	r, err := s.SubmitToBeatport(ctx, "r1")
	if err != nil || r.Status != domain.ReleaseSubmitted || r.BeatportID != "bp-1" {
		t.Fatalf("bad submit %+v %v", r, err)
	}
	if _, err := s.SubmitToBeatport(ctx, "r1"); !errors.Is(err, domain.ErrNotDraft) {
		t.Fatal("expected ErrNotDraft")
	}
	got, err := s.SplitSale(ctx, "r1", 1001)
	if err != nil || got["a"]+got["b"] != 1001 || got["b"] != 300 {
		t.Fatalf("bad split %v %v", got, err)
	}
}

func TestInvalidSplitsAndFailure(t *testing.T) {
	ctx := context.Background()
	s := newSvc(fakeBP{err: errors.New("boom")})
	if _, err := s.GenerateTrack(ctx, domain.CompositionRequest{CreatorID: "a"}, []domain.Split{{PayeeID: "a", BPS: 5000}}); !errors.Is(err, domain.ErrInvalidSplits) {
		t.Fatal("expected invalid splits")
	}
	s.GenerateTrack(ctx, domain.CompositionRequest{CreatorID: "a"}, nil)
	s.CreateRelease(ctx, "r1", "label", "EP", []string{"t1"})
	if r, err := s.SubmitToBeatport(ctx, "r1"); err == nil || r.Status != domain.ReleaseRejected {
		t.Fatal("expected rejected")
	}
}
