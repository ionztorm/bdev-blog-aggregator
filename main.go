package main

import (
	"fmt"
	"os"

	"github.com/ionztorm/gator/internal/command"
	"github.com/ionztorm/gator/internal/config"
	"github.com/ionztorm/gator/internal/database"
	"github.com/ionztorm/gator/internal/state"
	"github.com/ionztorm/gator/pkg/utils"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Failed to read config:", err)
		os.Exit(1)
	}

	db, err := database.ConnectToDB(cfg)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}
	defer utils.SafeClose(db.SQL)

	appState := &state.State{
		Cfg:    &cfg,
		DB:     db.Queries,
		DBConn: db.SQL,
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
