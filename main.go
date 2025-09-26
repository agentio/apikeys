package main

import (
	"os"

	"github.com/agentio/apikeys/cmd"
)

func main() {
	if err := cmd.Cmd().Execute(); err != nil {
		os.Exit(1)
	}
}
