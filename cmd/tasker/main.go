// Package main is Tasker entry point.
package main

import (
	"log"
	"os"

	"github.com/imarrche/tasker/internal/api"
	"github.com/imarrche/tasker/internal/config"
	"github.com/imarrche/tasker/internal/store/pg"
)

func main() {
	// Main logger.
	l := log.New(os.Stdout, "", log.LstdFlags)

	// Reading config from environment.
	c := config.New()

	// Opening PostgreSQL store.
	s := pg.New(c.PostgreSQL)
	if err := s.Open(); err != nil {
		l.Fatal(err)
	}

	// Starting the server.
	if err := api.NewServer(l, c, s).Start(); err != nil {
		l.Fatal(err)
	}
}
