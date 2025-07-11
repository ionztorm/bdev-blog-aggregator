package command

import (
	"context"
	"fmt"
	"gator/internal/rss"
	"gator/internal/state"
)

func init() {
	registerGlobal("agg", HandleAggregate)
}

func HandleAggregate(s *state.State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	result, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed from %s: %w", url, err)
	}

	fmt.Println(result)

	return nil
}
