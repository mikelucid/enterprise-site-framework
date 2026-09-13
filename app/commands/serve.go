package commands

import (
	"fmt"

	"github.com/mikelucid/enterprise-site-framework/bootstrap"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := bootstrap.SetupRouter()
		return r.Run(":8000")
	},
}

func init() { _ = fmt.Sprintf }
