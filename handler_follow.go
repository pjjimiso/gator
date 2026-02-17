package main

import (
	"context"
	"time"
	"fmt"

	"github.com/pkg/errors"
	"github.com/google/uuid"
	"github.com/pjjimiso/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("Follow command expects a single URL argument")
	}

	feedURL := cmd.args[0]
	feed, err := s.dbQueries.GetFeed(context.Background(), feedURL)
	if err != nil { 
		return errors.Wrap(err, "Failed to get feed")
	}

	followedFeed, err := s.dbQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	})
	if err != nil { 
		return errors.Wrap(err, "Failed to follow feed")
	}

	fmt.Printf("User %s is now following feed %s\n", followedFeed.UserName, followedFeed.FeedName)

	return nil
}
