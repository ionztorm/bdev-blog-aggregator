package main

import (
	"fmt"
	"gator/internal/command"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/state"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Failed to read config:", err)
		os.Exit(1)
	}

	_, dbQueries, err := database.ConnectToDB(cfg)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	appState := &state.State{
		Cfg: &cfg,
		DB:  dbQueries,
	}

	cmdRegistry := command.GetCmdRegistry()

	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	if err := cmdRegistry.Run(appState, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
