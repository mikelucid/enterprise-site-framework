package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("seed completed")
	},
}
