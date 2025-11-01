package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/marekbrze/gator/internal/database"
)

func registerHandler(s *state, cmd command) error {
	err := checkArguments(cmd, CheckArgumentsOptions{min: 1, max: 1})
	if err != nil {
		return err
	}
	userExists, err := s.db.UserExists(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if userExists {
		log.Fatal("User already exists")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.arguments[0]})
	if err != nil {
		return err
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error when setting the user in config. Details: %v", err)
	}
	fmt.Println("User has been created! \n\nUser details:")
	fmt.Println(user.Name, s.config.CurrentUserName)
	return nil
}
