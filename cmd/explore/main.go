package main

import (
	"fmt"
	"os"

	"github.com/willmadison/ossexplorer"
	"github.com/willmadison/ossexplorer/cli"
	"github.com/willmadison/ossexplorer/github"
	"github.com/willmadison/ossexplorer/mocks"
)

func main() {
	env := cli.Environment{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	var explorer ossexplorer.Explorer

	if accessToken == "" {
		fmt.Fprintf(env.Stdout, "warning: no GITHUB_ACCESS_TOKEN found defaulting to an in memory OSS explorer.\n")
		org := ossexplorer.Organization{Name: "InMemoryFakeOrg"}
		explorer = mocks.NewStubExplorer(org, nil)
	} else {
		explorer = github.NewExplorer(accessToken)
	}

	os.Exit(cli.Run(explorer, env))
}
