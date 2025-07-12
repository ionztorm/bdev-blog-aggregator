package command

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
	"log"
	"strconv"
)

func init() {
	registerGlobal("browse", middlewareLoggedIn(HandleBrowse))
}

func HandleBrowse(s *state.State, cmd Command, user database.User) error {

	limit := 2

	if len(cmd.Args) > 0 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = int(n)
		} else {
			log.Printf("Invalid limit provided: %v. Defaulting to %d.\n", err, limit)
		}
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts for %s: %w", user.Name, err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found.")
	} else {
		for _, post := range posts {
			fmt.Println("-------")
			fmt.Printf("Title: %s\n", post.Title)
			fmt.Printf("URL: %s\n", post.Url)
			if post.Description.Valid {
				fmt.Printf("Description: %s\n", post.Description.String)
			} else {
				fmt.Println("Description: (none)")
			}
			if post.PublishedAt.Valid {
				fmt.Printf("Published: %s\n", post.PublishedAt.Time.Format("2006-01-02 15:04:05"))
			} else {
				fmt.Println("Published: (unknown)")
			}
		}

	}

	return nil
}
