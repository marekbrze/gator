package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marekbrze/gator/internal/database"
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

func aggHandler(s *state, cmd command) error {
	err := checkArguments(cmd, 0)
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return &RSSFeed{}, fmt.Errorf("request to feed %s returned with status code %v - %v", feedURL, resp.StatusCode, resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return &RSSFeed{}, err
	}
	unescapeHTML(&rssFeed)
	return &rssFeed, nil
}

func unescapeHTML(rssFeed *RSSFeed) {
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
}

func addFeedHandler(s *state, cmd command, user database.User) error {
	err := checkArguments(cmd, 2)
	if err != nil {
		return err
	}
	newFeedDetails := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	}
	newFeed, err := s.db.CreateFeed(context.Background(), newFeedDetails)
	if err != nil {
		return err
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    newFeed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}
	fmt.Println(newFeed)
	return nil
}

func feedsHandler(s *state, cmd command) error {
	err := checkArguments(cmd, 0)
	if err != nil {
		return err
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("* %s (%s) added by %s\n", feed.Name, feed.Url, feed.AddedBy)
	}
	return nil
}

// func printRssFeed(rssFeed *RSSFeed) {
// 	fmt.Println(rssFeed.Channel.Title)
// 	fmt.Println(rssFeed.Channel.Description)
// 	for _, item := range rssFeed.Channel.Item {
// 		fmt.Println("* ", item.Title)
// 		fmt.Println("* ", item.Description)
// 	}
// }
