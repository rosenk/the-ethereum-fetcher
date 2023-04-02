package main

import (
	"os"

	"github.com/ju-popov/the-ethereum-fetcher/internal/cmd"
)

func main() {
	if err := cmd.NewLimeCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
