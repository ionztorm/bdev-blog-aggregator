package command

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"

	"github.com/google/uuid"
)

func init() {
	registerGlobal("addfeed", middlewareLoggedIn(HandleAddFeed))
}

func HandleAddFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("usage: addfeed <feedname> <url>")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feedId, now := database.GetCommonDBFields()

	userId := user.ID

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        feedId,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedName,
		Url:       feedURL,
		UserID:    userId,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    userId,
		FeedID:    feedId,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return fmt.Errorf("error following new feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
