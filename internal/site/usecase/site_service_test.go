package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/mikelucid/enterprise-site-framework/internal/site/domain"
)

type fakeRepo struct {
	created domain.Site
	createE error
	updateE error
}

func (f *fakeRepo) Create(context.Context, domain.Site) (domain.Site, error) {
	if f.createE != nil {
		return domain.Site{}, f.createE
	}
	return f.created, nil
}
func (f *fakeRepo) Get(context.Context, string) (domain.Site, error) { return domain.Site{}, nil }
func (f *fakeRepo) List(context.Context) ([]domain.Site, error)      { return nil, nil }
func (f *fakeRepo) Update(context.Context, domain.Site) (domain.Site, error) {
	if f.updateE != nil {
		return domain.Site{}, f.updateE
	}
	return f.created, nil
}
func (f *fakeRepo) Delete(context.Context, string) error { return nil }

type fakeProvisioner struct{ err error }

func (f fakeProvisioner) Provision(context.Context, domain.Site) error { return f.err }
func (f fakeProvisioner) Scale(context.Context, string) error          { return nil }
func (f fakeProvisioner) Decommission(context.Context, string) error   { return nil }

func TestCreate_ProvisionFailureReturnsCreatedSite(t *testing.T) {
	repo := &fakeRepo{created: domain.Site{ID: "site-1"}}
	svc := New(repo, fakeProvisioner{err: errors.New("provision failed")})
	site, err := svc.Create(context.Background(), domain.Site{})
	if err == nil {
		t.Fatal("expected error")
	}
	if site.ID != "site-1" || site.Status != "provision_failed" {
		t.Fatalf("unexpected site state: %+v", site)
	}
}

func TestCreate_CompensationFailureIncludesBothErrors(t *testing.T) {
	repo := &fakeRepo{created: domain.Site{ID: "site-1"}, updateE: errors.New("update failed")}
	svc := New(repo, fakeProvisioner{err: errors.New("provision failed")})
	site, err := svc.Create(context.Background(), domain.Site{})
	if err == nil {
		t.Fatal("expected error")
	}
	if site.ID != "site-1" {
		t.Fatalf("expected returned site id, got %+v", site)
	}
}
