package commands

import (
	"context"
	"fmt"

	"github.com/mikelucid/enterprise-site-framework/internal/agent"
	"github.com/spf13/cobra"
)

var agentRunCmd = &cobra.Command{
	Use:   "agent:run",
	Short: "Run the framework automation agent runtime",
	RunE: func(cmd *cobra.Command, args []string) error {
		runner := agent.NewRunner()
		fmt.Println("agent runtime started")
		if err := runner.Run(context.Background()); err != nil {
			return err
		}
		return nil
	},
}
