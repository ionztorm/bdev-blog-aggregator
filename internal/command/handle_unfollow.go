package command

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
)

func init() {
	registerGlobal("unfollow", middlewareLoggedIn(HandleUnfollow))
}

func HandleUnfollow(s *state.State, cmd Command, user database.User) error {

	userId := user.ID

	if len(cmd.Args) < 1 {
		return errors.New("usage: unfollow <feed url>")
	}

	feedURL := cmd.Args[0]

	feed, err := s.DB.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error locating specified feed: %w", err)
	}

	feedId := feed.ID

	err = s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: userId,
		FeedID: feedId,
	})

	if err != nil {
		return fmt.Errorf("error unfollowing feeed: %w", err)
	}

	fmt.Printf("%s successfully unfollowed for %s", feedURL, user.Name)

	return nil
}
