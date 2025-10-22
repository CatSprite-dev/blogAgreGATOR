package main

import (
	"log"
	"os"

	"github.com/CatSprite-dev/blogAgreGATOR/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state{
		cfg: &cfg,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("command required")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.run(&programState, cmd); err != nil {
		log.Fatal(err)
	}
}
