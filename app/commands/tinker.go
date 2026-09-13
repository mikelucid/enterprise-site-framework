package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mikelucid/enterprise-site-framework/app/facades"
	"github.com/spf13/cobra"
)

var tinkerCmd = &cobra.Command{
	Use:   "tinker",
	Short: "Interactive shell",
	Run: func(cmd *cobra.Command, args []string) {
		repl()
	},
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	vars := map[string]string{}
	history := []string{}
	fmt.Println("Tinker shell (type 'help' for commands, 'exit' to quit)")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		history = append(history, line)
		switch {
		case line == "exit":
			return
		case line == "help":
			fmt.Println("Commands: help, history, set <k>=<v>, get <k>, Site.Create(), Character.Find(<id>)")
		case line == "history":
			for _, h := range history {
				fmt.Println(h)
			}
		case strings.HasPrefix(line, "set "):
			kv := strings.TrimPrefix(line, "set ")
			parts := strings.SplitN(kv, "=", 2)
			if len(parts) != 2 {
				fmt.Println("error: invalid assignment")
				continue
			}
			vars[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			fmt.Println("ok")
		case strings.HasPrefix(line, "get "):
			k := strings.TrimSpace(strings.TrimPrefix(line, "get "))
			fmt.Println(vars[k])
		case strings.HasPrefix(line, "Site.Create"):
			fmt.Println(facades.Site().Create(map[string]any{"name": "interactive-site"}))
		case strings.HasPrefix(line, "Character.Find"):
			fmt.Println(facades.Character().Find("interactive-character"))
		default:
			fmt.Println("error: unknown command")
		}
	}
}
