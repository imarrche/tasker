// Package main is Tasker entry point.
package main

import (
	"github.com/imarrche/tasker/internal/api"
	"github.com/imarrche/tasker/internal/store/inmem"
)

func main() {
	api.NewServer(inmem.TestStoreWithFixtures()).Start()
}
