package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
	"gator/pkg/utils"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error making request: %w", err)
	}
	defer utils.SafeClose(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var rssFeed RSSFeed

	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return nil, fmt.Errorf("something got messed up: %w", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return &rssFeed, nil
}

func ScrapeFeeds(s *state.State) {
	nextFeed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err == sql.ErrNoRows {
		fmt.Println("No feeds available to fetch.")
		return
	} else if err != nil {
		fmt.Printf("Error fetching next feed: %v\n", err)
		return
	}

	now := time.Now()

	err = s.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            nextFeed.ID,
		UpdatedAt:     now,
		LastFetchedAt: sql.NullTime{Time: now, Valid: true},
	})

	if err != nil {
		fmt.Println("There was a problem marking the feed as fetched.")
		return
	}

	feedContent, err := FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Printf("error fetching feed content: %v", err)
		return
	}

	items := feedContent.Channel.Item
	separator := "================================"
	fmt.Println(separator)
	fmt.Printf("Scraping feed: %s (%s)\n", nextFeed.Name, nextFeed.Url)
	fmt.Printf("Found %d items in the feed:\n", len(items))
	fmt.Println(separator)
	fmt.Println()

	for _, item := range items {
		fmt.Println(" * ", item.Title)
	}
}
