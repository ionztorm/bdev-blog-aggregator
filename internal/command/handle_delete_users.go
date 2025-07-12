package command

import (
	"context"
	"fmt"

	"github.com/ionztorm/gator/internal/state"
)

func init() {
	registerGlobal("reset", HandleDeleteUsers)
}

func HandleDeleteUsers(s *state.State, cmd Command) error {

	err := s.DB.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users: %w", err)
	}

	fmt.Println("all users were successfully deleted. You will need to register again.")

	return nil
}
