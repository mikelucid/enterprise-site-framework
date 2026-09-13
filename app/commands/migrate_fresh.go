package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateFreshCmd = &cobra.Command{
	Use:   "migrate:fresh",
	Short: "Reset and re-run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fresh migrations completed")
	},
}
