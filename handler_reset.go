package main

import (
	"context"
	
	"github.com/pkg/errors"
)


func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("Reset command doesn't expect any arguments")
	}

	err := s.dbQueries.TruncateUsers(context.Background())
	if err != nil { 
		return errors.Wrap(err, "User reset failed")
	}

	return nil
}
