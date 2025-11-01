package main

import (
	"context"
	"fmt"

	"github.com/marekbrze/gator/internal/database"
)

func usersHandler(s *state, cmd command) error {
	err := checkArguments(cmd, CheckArgumentsOptions{min: 0, max: 0})
	if err != nil {
		return err
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error when getting users from the database: %s", err)
	}
	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUserInfo, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, currentUserInfo)
	}
}
