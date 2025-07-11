package command

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
	"time"

	"github.com/google/uuid"
)

func init() {
	registerGlobal("addfeed", HandleAddFeed)
}

func HandleAddFeed(s *state.State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("usage: addfeed <feedname> <url>")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	uuid := uuid.New()
	now := time.Now()

	currentUser := s.Cfg.CurrentUserName

	user, err := s.DB.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("unable to locate user '%s': %w", currentUser, err)
	}

	userId := user.ID

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedName,
		Url:       feedURL,
		UserID:    userId,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
