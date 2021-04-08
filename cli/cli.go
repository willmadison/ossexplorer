package cli

import (
	"context"
	"fmt"
	"io"
	"text/tabwriter"

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
	Org        string `arg required help:"The organization to explore/search."`
	Sort       string `help:"The axis by which our exploratory results should be sorted (stars|forks|pulls|contrib_rate). (default=stars)" default:"stars"`
	Order      string `help:"The orientation of which our exploratory results should be sorted (desc|asc). (default=desc)" default:"desc"`
	MaxResults int    `help:"The maximum number of results we'd like to limit our display to. (default=10)" default:"10"`
}

func (cmd *ReposCmd) Run(ctx *Context) error {
	cntx := context.Background()
	organization, err := ctx.explorer.FindOrganization(cntx, cmd.Org)

	if err != nil {
		fmt.Fprintf(ctx.env.Stderr, "Unable to locate the given organization (%v). Please ensure you have access and that it exists.\n", cmd.Org)
		return nil
	}

	var modifier ossexplorer.RepositoryResultModifier

	switch cmd.Sort + "_" + cmd.Order {
	case "stars_asc":
		modifier = ossexplorer.ByStarsAscending
	case "stars_desc":
		modifier = ossexplorer.ByStarsDescending
	case "forks_asc":
		modifier = ossexplorer.ByForksAscending
	case "forks_desc":
		modifier = ossexplorer.ByForksDescending
	case "pulls_asc":
		modifier = ossexplorer.ByPullRequestsAscending
	case "pulls_desc":
		modifier = ossexplorer.ByPullRequestsDescending
	case "contrib_rate_asc":
		modifier = ossexplorer.ByContributionRateAscending
	case "contrib_rate_desc":
		modifier = ossexplorer.ByContributionRateDescending
	}

	repos, err := ctx.explorer.FindRepositoriesFor(cntx, organization, modifier)

	if err != nil {
		fmt.Fprintf(ctx.env.Stderr, "Repository lookup failed for %v. Please ensure you have read access to this organization's repositories.\n", cmd.Org)
		return nil
	}

	if len(repos) == 0 {
		fmt.Fprintf(ctx.env.Stdout, "No repositories found for %v. Nothing to do here but chill 😎...\n", cmd.Org)
		return nil
	}

	w := tabwriter.NewWriter(ctx.env.Stdout, 0, 0, 1, ' ', tabwriter.DiscardEmptyColumns|tabwriter.AlignRight)

	fmt.Fprintln(w, "#\tRepository\t# Stars\t# Forks\t# Pull Requests\tContribution %\t")
	fmt.Fprintln(w, "_\t__________\t_______\t_______\t_______________\t______________\t")

	for i, repo := range repos {
		if i == cmd.MaxResults {
			break
		}

		fmt.Fprintf(w, "%d\t%v\t%v\t%v\t%v\t%.2f%%\t\n", i+1, repo.Name, repo.Stars, repo.Forks, repo.PullRequests, repo.ContributionRate()*100)
	}

	w.Flush()

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
