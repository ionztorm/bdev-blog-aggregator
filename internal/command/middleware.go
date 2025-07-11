package command

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
)

func middlewareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
	return func(s *state.State, cmd Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user is not logged in: %w", err)
		}

		return handler(s, cmd, user)
	}

}
