package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/mikelucid/enterprise-site-framework/internal/site/domain"
)

type Service struct {
	repo        domain.Repository
	provisioner domain.Provisioner
}

func New(repo domain.Repository, provisioner domain.Provisioner) *Service {
	return &Service{repo: repo, provisioner: provisioner}
}

func (s *Service) Create(ctx context.Context, site domain.Site) (domain.Site, error) {
	created, err := s.repo.Create(ctx, site)
	if err != nil {
		return domain.Site{}, err
	}
	if err := s.provisioner.Provision(ctx, created); err != nil {
		created.Status = "provision_failed"
		if _, updateErr := s.repo.Update(ctx, created); updateErr != nil {
			return created, fmt.Errorf("provision failed and compensation update failed: %w", errors.Join(err, updateErr))
		}
		return created, err
	}
	return created, nil
}

func (s *Service) Get(ctx context.Context, id string) (domain.Site, error) { return s.repo.Get(ctx, id) }
