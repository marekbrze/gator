package main

import (
	"github.com/marekbrze/gator/internal/config"
	"github.com/marekbrze/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
