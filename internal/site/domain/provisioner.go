package domain

import "context"

type Provisioner interface {
	Provision(context.Context, Site) error
	Scale(context.Context, string) error
	Decommission(context.Context, string) error
}
