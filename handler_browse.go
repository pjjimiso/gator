package main

import (
	"context"
	"strconv"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pjjimiso/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.args) > 1 {
		errors.New("usage: browse [limit]")
	}
	if len(cmd.args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return errors.Wrap(err, "Invalid limit argument, must be a number")
		}
	}

	posts, err := s.dbQueries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:	int32(limit),
	})
	if err != nil { 
		return errors.Wrapf(err, "Couldn't get posts for user %s", user.Name)
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
