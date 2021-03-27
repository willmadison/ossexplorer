package cli

import (
	"fmt"
	"io"

	"github.com/alecthomas/kong"
	"github.com/willmadison/ossexplorer"
)

// Environment provides an abstraction around the execution environment
type Environment struct {
	Stderr io.Writer
	Stdout io.Writer
	Stdin  io.Reader
}

type Context struct {
	explorer ossexplorer.Explorer
	env      Environment
}

type ReposCmd struct {
	Org   string `arg required help:"The organization to explore/search."`
	Sort  string `arg help:"The axis by which our exploratory results should be sorted (stars|forks|pulls|contrib_rate)." default:"stars"`
	Order string `arg help:"The orientation of which our exploratory results should be sorted (desc|asc)." default:"desc"`
}

func (cmd *ReposCmd) Run(ctx *Context) error {
	fmt.Println("Org =", cmd.Org)
	fmt.Println("Sort =", cmd.Sort)
	fmt.Println("Order =", cmd.Order)

	return nil
}

type CLI struct {
	Context
	Repos ReposCmd `cmd help:"Searches/Explores OSS Repos along various axes."`
}

func Run(explorer ossexplorer.Explorer, env Environment) int {
	app := CLI{
		Context: Context{
			explorer: explorer,
			env:      env,
		},
	}

	ctx := kong.Parse(&app,
		kong.Name("explore"),
		kong.Description("An awesome OSS explorer"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)
	err := ctx.Run(&app.Context)
	ctx.FatalIfErrorf(err)

	return 0
}
