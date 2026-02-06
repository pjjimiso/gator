package main


import (
	"os"
	"log"

	"github.com/pjjimiso/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct { 
	name string
	args []string
}


func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide a command argument")
	}

	configFile, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := state{ cfg: &configFile }

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	cmd := command {
		name: commandName,
		args: commandArgs,
	}

	c := commands{
		cmdMap: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)

	err = c.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
