package main 

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pjjimiso/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.dbQueries.GetFollowedFeeds(context.Background(), user.Name)
	if err != nil {
		return errors.Wrap(err, "Failed to get followed feeds")
	}

	fmt.Println("You currently follow these feeds:")
	for _, feed := range feeds { 
		fmt.Printf(" - %s\n", feed.FeedName)
	}

	return nil	
}
