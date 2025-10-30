package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.list[cmd.name]
	if !exists {
		return fmt.Errorf("handler for command %v doesn't exists", cmd.name)
	}
	err := handler(s, cmd)
	if err != nil {
		return fmt.Errorf("there was a problem when running a %v command: %v", cmd.name, err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}

func checkArguments(cmd command, limit int) error {
	if len(cmd.arguments) < limit {
		return fmt.Errorf("number of required arguments: %d", limit)
	}

	if len(cmd.arguments) > limit {
		return fmt.Errorf("too many arguments. Number of required arguments: %d", limit)
	}
	return nil
}
