package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/marekbrze/gator/internal/database"
)

func browseHandler(s *state, cmd command, user database.User) error {
	err := checkArguments(cmd, CheckArgumentsOptions{min: 0, max: 1})
	if err != nil {
		return err
	}
	limit := 2
	if len(cmd.arguments) > 0 {

		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}
	}
	posts, err := s.db.GetPostsForTheUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	if limit > len(posts) {
		limit = len(posts)
	}
	for i, post := range posts[:limit] {
		fmt.Printf("%d. (%s) %s\n", i+1, post.Feedname, post.Title)
	}
	return nil
}
