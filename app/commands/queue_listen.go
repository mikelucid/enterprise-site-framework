package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var queueListenCmd = &cobra.Command{
	Use:   "queue:listen",
	Short: "Start queue listener",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("queue listener started")
	},
}
