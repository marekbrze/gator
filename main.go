package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marekbrze/gator/internal/config"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	gatorState := state{config: &configFile}

	gatorCmds := commands{make(map[string]func(*state, command) error)}
	gatorCmds.register("login", loginHandler)
	args := os.Args
	if len(args) <= 1 {
		log.Fatal("you need to provide arguments to the gator program")
	}
	command := command{
		name:      args[1],
		arguments: args[2:],
	}
	err = gatorCmds.run(&gatorState, command)
	if err != nil {
		log.Fatal(err)
	}
}
