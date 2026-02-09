package main 

import ( 
	"fmt"
	"time"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pjjimiso/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error { 
	if len(cmd.args) != 1 {
		return errors.New("Expecting a single user argument")

	}
	if cmd.args[0] == "" {
		return errors.New("Username is invalid")
	}

	name := cmd.args[0]
	user, err := s.dbQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		name,
	})
	if err != nil { 
		return errors.Wrap(err, "User creation failed")
	}

	err = s.cfg.SetUser(name)
	if err != nil { 
		return errors.Wrap(err, "Failed to set user in config")
	}

	fmt.Printf("User created:\n name: %s\n created_at: %v\n updated_at: %v\n id: %s\n",
		user.Name, user.CreatedAt, user.UpdatedAt, user.ID)
	
	return nil
}

