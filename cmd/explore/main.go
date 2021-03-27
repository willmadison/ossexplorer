package main

import (
	"os"

	"github.com/willmadison/ossexplorer/cli"
	"github.com/willmadison/ossexplorer/github"
)

func main() {
	env := cli.Environment{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	if accessToken == "" {
		panic("missing GITHUB_ACCESS_TOKEN. Please ensure its appropriately set in your environment.")
	}

	explorer := github.NewExplorer(accessToken)
	os.Exit(cli.Run(explorer, env))
}
