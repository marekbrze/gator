package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	err := checkArguments(cmd, CheckArgumentsOptions{min: 0, max: 1})
	if err != nil {
		return err
	}
	err = s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users table")
	}
	fmt.Println("Users table has been emptied")
	return nil
}
