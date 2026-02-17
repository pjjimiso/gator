package main

import ( 
	"fmt"
	"time"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pjjimiso/gator/internal/database"
)


func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("Addfeed expects 2 arguments: feedName and feedURL")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	feed, err := s.dbQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		feedName,
		Url:		feedURL,
		UserID:		user.ID,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to create feed")
	}

	_, err = s.dbQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to follow feed")
	}

	printFeedDB(feed)

	return nil
}

func printFeedDB(feed database.Feed) {
	fmt.Printf("ID:		%s\n", feed.ID)
	fmt.Printf("CreatedAt:	%s\n", feed.CreatedAt)
	fmt.Printf("UpdatedAt:	%s\n", feed.UpdatedAt)
	fmt.Printf("Name:	%s\n", feed.UpdatedAt)
	fmt.Printf("Url:	%s\n", feed.Url)
	fmt.Printf("UserID:	%s\n", feed.UserID)
}
