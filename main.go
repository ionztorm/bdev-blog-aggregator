package main

import (
	"fmt"
	"gator/internal/command"
	"gator/internal/config"
	"os"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	state := &command.State{
		Cfg: &cfg,
	}

	cmdRegistry := command.Commands{
		Handlers: make(map[string]func(*command.State, command.Command) error),
	}

	cmdRegistry.Register("login", command.HandleLogin)

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

	if err := cmdRegistry.Run(state, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
