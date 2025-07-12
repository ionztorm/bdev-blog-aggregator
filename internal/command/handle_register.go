package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ionztorm/gator/internal/database"
	"github.com/ionztorm/gator/internal/state"

	"github.com/google/uuid"
)

func init() {
	registerGlobal("register", HandleRegister)
}

func HandleRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register requires a username argument")
	}

	username := cmd.Args[0]
	ctx := context.Background()

	_, err := s.DB.GetUser(ctx, username)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking for user existence: %w", err)
	}
	if err == nil {
		return fmt.Errorf("user %q already exists", username)
	}

	id := uuid.New()
	now := time.Now()

	_, err = s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("User %q created successfully!\n", username)

	return nil
}
