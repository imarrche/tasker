// Package main is Tasker entry point.
package main

import "github.com/imarrche/tasker/internal/api"

func main() {
	api.NewServer().Start()
}
