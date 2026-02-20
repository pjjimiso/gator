package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pjjimiso/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("usage: unfollow <feed_url>")
	}

	feedUrl := cmd.args[0]
	err := s.dbQueries.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Url:	feedUrl,
		Name:	user.Name,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to unfollow feed")
	}

	return nil
}
