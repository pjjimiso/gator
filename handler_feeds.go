package main

import (
	"fmt"
	"context"

	"github.com/pkg/errors"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("Feeds command doesn't expect an arguments")
	}

	feeds, err := s.dbQueries.GetFeeds(context.Background())
	if err != nil {
		return errors.Wrap(err, "Failed to get feeds")
	}

	for _, feed := range feeds { 
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("Username: %s\n\n", feed.Username)
	}

	return nil
}
