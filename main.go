package main


import (
	"os"
	"log"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pjjimiso/gator/internal/config"
	"github.com/pjjimiso/gator/internal/database"
)

type state struct {
	cfg		*config.Config
	dbQueries	*database.Queries
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide a command argument")
		os.Exit(1)
	}
	if os.Args[1] == "" {
		log.Fatal("Command cannot be empty")
		os.Exit(1)
	}

	configFile, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	programState := &state{ cfg: &configFile }
	db, err := sql.Open("postgres", programState.cfg.DbURL)
	programState.dbQueries = database.New(db)

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{
		name: cmdName,
		args: cmdArgs,
	})
	if err != nil { 
		log.Fatal(err)
	}
}
