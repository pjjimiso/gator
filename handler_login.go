package main

import (
	"context"

	"github.com/pkg/errors"
)


func handlerLogin(s *state, cmd command) error { 
	if len(cmd.args) != 1 {
		return errors.New("Expecting a single user argument")
	}
	if cmd.args[0] == "" { 
		return errors.New("Username is invalid")
	}

	username := cmd.args[0]

	_, err := s.dbQueries.GetUser(context.Background(), username)
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve user")
	}

	err = s.cfg.SetUser(username)
	if err != nil { 
		return errors.Wrap(err, "Failed to set user in config")
	}

	return nil
}
