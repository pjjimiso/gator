package main

import ( 
	"fmt"
	"time"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pjjimiso/gator/internal/database"
)


func handlerFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("Addfeed expects 2 arguments: feedName and feedURL")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]
	currentUser, err := s.dbQueries.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errors.Wrap(err, "Failed to get user")
	}

	feed, err := s.dbQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		feedName,
		Url:		feedURL,
		UserID:		currentUser.ID,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to create feed")
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
