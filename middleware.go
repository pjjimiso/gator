package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pjjimiso/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.dbQueries.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return errors.Wrap(err, "Failed to get user")
		}
		return handler(s, cmd, user)
	}
}
