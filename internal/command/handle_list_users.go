package command

import (
	"context"
	"fmt"

	"github.com/ionztorm/gator/internal/state"
)

func init() {
	registerGlobal("users", HandleListUsers)
}

func HandleListUsers(s *state.State, cmd Command) error {
	users, err := s.DB.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving users: %w", err)
	}

	for _, user := range users {
		username := user
		if username == s.Cfg.CurrentUserName {
			username = fmt.Sprintf("%v (current)", username)
		}
		fmt.Printf("* %s\n", username)
	}

	return nil

}
