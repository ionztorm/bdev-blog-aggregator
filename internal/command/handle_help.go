package command

import (
	"fmt"
	"strings"

	"gator/internal/state"
)

var helpNotes = map[string]string{
	"help":      "gator help [command]         - Show help info",
	"migrate":   "gator migrate [up|down]      - Run DB migrations",
	"register":  "gator register <name>        - Create a user",
	"login":     "gator login <name>           - Log in",
	"addfeed":   "gator addfeed <url>          - Add an RSS feed (auto-follows it)",
	"feeds":     "gator feeds                  - List all feeds added by any user",
	"follow":    "gator follow <url>           - Follow a feed",
	"following": "gator following              - List feeds you’re currently following",
	"unfollow":  "gator unfollow <url>         - Unfollow a feed",
	"agg":       "gator agg <interval>         - Start the aggregator loop (e.g., 5s, 1m, 1h)",
	"browse":    "gator browse <limit>         - Browse posts from followed feeds (optional limit)",
	"list":      "gator list users             - List all registered users",
	"reset":     "gator reset                  - Delete all users and cascade-delete related data",
}

func init() {
	registerGlobal("help", HandleHelp)
}

func HandleHelp(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		fmt.Println("Available commands:")
		fmt.Println()
		for _, desc := range helpNotes {
			fmt.Println("  " + desc)
		}
		return nil
	}

	query := strings.ToLower(cmd.Args[0])
	desc, ok := helpNotes[query]
	if !ok {
		return fmt.Errorf("unknown command %q. Use `gator help` to list all commands", query)
	}

	fmt.Printf("Help for `%s`:\n\n  %s\n", query, desc)
	return nil
}
