package main

import "fmt"

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.cmdMap[cmd.name]
	if !ok {
		return fmt.Errorf("That command does not exist")
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func(c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}
