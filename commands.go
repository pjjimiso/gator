package main

import (
	"github.com/pkg/errors"
)

type command struct { 
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	registeredCommand, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("That command does not exist")
	}

	err := registeredCommand(s, cmd)
	if err != nil {
		return errors.Wrap(err, "command failed")
	}
	return nil
}

func(c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
