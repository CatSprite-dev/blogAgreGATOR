package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/CatSprite-dev/blogAgreGATOR/internal/config"
	"github.com/CatSprite-dev/blogAgreGATOR/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("Not connect to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerFeeds)

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
	_, err = fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		fmt.Println("fetching error")
	}
}
