package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func loginHandler(s *state, cmd command) error {
	err := checkArguments(cmd, 1)
	if err != nil {
		return err
	}
	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err == sql.ErrNoRows {
		log.Fatal("User doesn't exist")
	}
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error when setting the user in config. Details: %v", err)
	}
	fmt.Println("user has been set. Username:", s.config.CurrentUserName)
	return nil
}
