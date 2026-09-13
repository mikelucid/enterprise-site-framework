package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "framework",
	Short: "NOVA enterprise site framework",
}

func Execute() {
	rootCmd.AddCommand(serveCmd, migrateCmd, migrateFreshCmd, queueListenCmd, seedCmd, tinkerCmd, agentRunCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
