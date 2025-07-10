package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
	"time"

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
	name := sql.NullString{
		String: username,
		Valid:  true,
	}

	_, err := s.DB.GetUser(ctx, name)
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
		Name:      name,
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
