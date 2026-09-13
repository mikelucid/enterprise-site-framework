package adapter

import (
	"context"

	"github.com/mikelucid/enterprise-site-framework/internal/site/domain"
)

type TerraformProvisioner struct{}

func (TerraformProvisioner) Provision(context.Context, domain.Site) error { return nil }
func (TerraformProvisioner) Scale(context.Context, string) error          { return nil }
func (TerraformProvisioner) Decommission(context.Context, string) error   { return nil }
