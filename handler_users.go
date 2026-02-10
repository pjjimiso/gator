package main

import (
	"fmt"
	"context"

	"github.com/pkg/errors"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		errors.New("Users command doesn't expect any arguments")
	}

	users, err := s.dbQueries.ListUsers(context.Background())
	if err != nil { 
		errors.Wrap(err, "Failed to get users")
	}

	for _, user := range users {
		line := fmt.Sprintf(" * %s ", user)
		if user == s.cfg.Username { 
			line += "(current)"
		}
		line += "\n"
		fmt.Printf(line)
	}

	return nil
}
