package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gator/internal/state"
)

func init() {
	registerGlobal("login", HandleLogin)
}

func HandleLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires a username argument")
	}

	username := cmd.Args[0]

	ctx := context.Background()
	name := sql.NullString{
		String: username,
		Valid:  true,
	}

	_, err := s.DB.GetUser(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user %q does not exist", username)
		}
		return fmt.Errorf("error checking for user existence: %w", err)
	}
	err = s.Cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User successfully updated to %s", username)

	return nil
}
