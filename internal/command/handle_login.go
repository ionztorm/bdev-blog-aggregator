package command

import (
	"errors"
	"fmt"
)

func HandleLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires a username argument")
	}

	username := cmd.Args[0]

	err := s.Cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User successfully updated to %s", username)

	return nil
}
