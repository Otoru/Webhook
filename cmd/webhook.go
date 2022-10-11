package main

import (
	"os"

	"github.com/otoru/webhook/pkg/cmd"
)

func main() {
	instance := cmd.CreateRootCommand(os.Stdout)
	if err := instance.Execute(); err != nil {
		os.Exit(1)
	}
}
