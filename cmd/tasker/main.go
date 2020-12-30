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
	l := log.New(os.Stdout, "", log.LstdFlags)

	c := config.New() // Reading config from environment.

	// PostgreSQL.
	store := pg.New(c.PostgreSQL)
	if err := store.Open(); err != nil {
		l.Fatal(err)
	}

	api.NewServer(c, l, store).Start()
}
