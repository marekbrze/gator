package main

import "fmt"

func loginHandler(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("loginHandler functions requires username argument")
	}
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("too many arguments. Login function requires only 1 argument: username")
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error when setting the user in config. Details: %v", err)
	}
	fmt.Println("user has been set. Username:", cmd.arguments[0])
	return nil
}
