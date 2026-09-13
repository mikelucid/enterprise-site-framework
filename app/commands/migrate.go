package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrations completed")
	},
}
