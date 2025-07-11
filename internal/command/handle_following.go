package command

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
)

func init() {
	registerGlobal("following", middlewareLoggedIn(HandleFollowing))
}

func HandleFollowing(s *state.State, cmd Command, user database.User) error {
	username := s.Cfg.CurrentUserName

	user_id := user.ID

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user_id)
	if err != nil {
		return fmt.Errorf("error looking up follows for user %s: %w", username, err)
	}

	if len(feeds) < 1 {
		fmt.Println("You are not currently following any feeds.")
		return nil
	}

	fmt.Println("Your followed feeds:")
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}
