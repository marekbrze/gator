package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/marekbrze/gator/internal/config"
	"github.com/marekbrze/gator/internal/database"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", configFile.DBURL)
	if err != nil {
		log.Fatal("Couldn't connect to the database")
	}
	dbQueries := database.New(db)

	gatorState := state{config: &configFile, db: dbQueries}
	gatorCmds := commands{make(map[string]func(*state, command) error)}
	// Users commands
	gatorCmds.register("login", loginHandler)
	gatorCmds.register("register", registerHandler)
	gatorCmds.register("users", usersHandler)
	// Feeds commands
	gatorCmds.register("agg", aggHandler)
	gatorCmds.register("feeds", feedsHandler)
	gatorCmds.register("addfeed", middlewareLoggedIn(addFeedHandler))
	// Follow commands
	gatorCmds.register("follow", middlewareLoggedIn(followHandler))
	gatorCmds.register("following", middlewareLoggedIn(followingHandler))
	gatorCmds.register("unfollow", middlewareLoggedIn(unfollowHandler))
	// Reset
	gatorCmds.register("reset", resetHandler)
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
