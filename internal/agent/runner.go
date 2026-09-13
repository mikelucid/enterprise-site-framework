package agent

import "context"

type Runner struct{}

func NewRunner() *Runner { return &Runner{} }

func (r *Runner) Run(context.Context) error { return nil }
