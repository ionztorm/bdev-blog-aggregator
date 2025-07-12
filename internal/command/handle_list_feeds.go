package command

import (
	"context"
	"fmt"

	"github.com/ionztorm/gator/internal/state"
)

func init() {
	registerGlobal("feeds", HandleListFeeds)
}

func HandleListFeeds(s *state.State, cmd Command) error {

	feeds, err := s.DB.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("unable to retrieve feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* Name: %s, URL: %s, User: %s\n", feed.Name, feed.Url, feed.User.String)
	}

	return nil
}
