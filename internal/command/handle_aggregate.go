package command

import (
	"errors"
	"fmt"
	"gator/internal/aggregate"
	"gator/internal/state"
	"time"
)

func init() {
	registerGlobal("agg", HandleAggregate)
}

func HandleAggregate(s *state.State, cmd Command) error {

	if len(cmd.Args) < 1 {
		return errors.New("usage: agg <interval>")
	}

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		aggregate.ScrapeFeeds(s)
	}
}
