package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/marekbrze/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}
	feedArticles, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, article := range feedArticles.Channel.Item {
		layout := time.RFC1123Z
		pubDate, err := time.Parse(layout, article.PubDate)
		if err != nil {
			return err
		}
		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       article.Title,
			Url:         article.Link,
			Description: article.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}
		post, err := s.db.CreatePost(context.Background(), postParams)
		duplicateRowError := "duplicate key value"
		if err != nil {
			if strings.Contains(err.Error(), duplicateRowError) {
				continue
			}
			return err
		}
		fmt.Println(post.Title)
	}

	return nil
}
