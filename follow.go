package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/marekbrze/gator/internal/database"
)

func followHandler(s *state, cmd command, user database.User) error {
	err := checkArguments(cmd, 1)
	if err != nil {
		return err
	}

	feedID, err := s.db.FindFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}
	fmt.Printf("Feed Name: %s\nUser: %s", feedFollow.FeedName, feedFollow.UserName)
	return nil
}

func followingHandler(s *state, cmd command, user database.User) error {
	err := checkArguments(cmd, 0)
	if err != nil {
		return err
	}

	feedsFollowedByUser, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, feed := range feedsFollowedByUser {
		fmt.Printf("- %s", feed.FeedName)
	}
	return nil
}

func unfollowHandler(s *state, cmd command, user database.User) error {
	err := checkArguments(cmd, 1)
	if err != nil {
		return err
	}

	feedID, err := s.db.FindFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	feedFollowParams := database.DeleteFeedFollowForUserParams{
		UserID: user.ID,
		FeedID: feedID,
	}
	err = s.db.DeleteFeedFollowForUser(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}
	fmt.Printf("Feed %s has been deleted for user: %s", cmd.arguments[0], user.Name)
	return nil
}
