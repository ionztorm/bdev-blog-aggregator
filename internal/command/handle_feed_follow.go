package command

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
)

func init() {
	registerGlobal("follow", middlewareLoggedIn(HandleFeedFollow))
}

func HandleFeedFollow(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("usage: follow <feed url>")
	}

	feedURL := cmd.Args[0]
	feed, err := s.DB.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error looking up feed: %w", err)
	}

	feed_id := feed.ID

	id, now := database.GetCommonDBFields()

	user_id := user.ID

	result, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user_id,
		FeedID:    feed_id,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Printf("User %s is now following %s", result.UserName, result.FeedName)

	return nil
}
