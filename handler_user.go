package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/CatSprite-dev/blogAgreGATOR/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command expects a single argument, the username")
	}
	userName := cmd.Args[0]
	if _, err := s.db.GetUserByName(context.Background(), userName); err != nil {
		fmt.Println("User not registered")
		os.Exit(1)
	}

	err := s.cfg.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("register command expects a single argument, the username\nusage: %v <name>", cmd.Name)
	}
	userID := uuid.New()
	userName := cmd.Args[0]
	if _, err := s.db.GetUserByName(context.Background(), userName); err == nil {
		fmt.Println("User with that name already exists")
		os.Exit(1)
	}

	userParams := database.CreateUserParams{
		ID:        userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	newUser, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("user not create: %v", err)
	}

	err = s.cfg.SetUser(newUser.Name)
	if err != nil {
		return err
	}

	fmt.Println("User was created")
	fmt.Printf("User id: %v\n", newUser.ID)
	fmt.Printf("User name: %v\n", newUser.Name)
	fmt.Printf("User creted at: %v\n", newUser.CreatedAt)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("error reseting table: %v", err)
	}
	fmt.Println("Reset successful")
	os.Exit(0)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users %v", err)
	}
	if len(users) == 0 {
		fmt.Println("No users have registered yet")
		fmt.Println("To register, please use")
		fmt.Println("register <name>")
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("%v (current)\n", user.Name)
		} else {
			fmt.Printf("%v\n", user.Name)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
