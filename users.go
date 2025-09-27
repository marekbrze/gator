package main

import (
	"context"
	"fmt"
)

func usersHandler(s *state, cmd command) error {
	err := checkArguments(cmd, 0)
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
