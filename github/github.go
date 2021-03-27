package github

import (
	"context"

	"github.com/willmadison/ossexplorer"

	gh "github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
)

type GitHubExplorer struct {
	client *gh.Client
}

func (g *GitHubExplorer) FindOrganization(ctx context.Context, name string) (ossexplorer.Organization, error) {
	org, _, err := g.client.Organizations.Get(ctx, name)

	if err != nil {
		return ossexplorer.Organization{}, err
	}

	return ossexplorer.Organization{
		Name: *org.Login,
	}, nil
}

func (g *GitHubExplorer) FindRepositoriesFor(ctx context.Context, org ossexplorer.Organization, modifiers ...ossexplorer.RepositoryResultModifier) ([]ossexplorer.Repository, error) {
	options := &gh.RepositoryListByOrgOptions{
		ListOptions: gh.ListOptions{
			PerPage: 10,
		},
	}

	var repositories []ossexplorer.Repository

	var allRepos []*gh.Repository

	var done bool

	for !done {
		repos, resp, err := g.client.Repositories.ListByOrg(ctx, org.Name, options)

		if err != nil {
			return repositories, err
		}

		allRepos = append(allRepos, repos...)

		done = resp.NextPage == 0
		options.Page = resp.NextPage
	}

	for _, r := range allRepos {
		pulls, err := g.getPullRequestCount(ctx, org.Name, r.GetName())

		if err != nil {
			return repositories, err
		}

		repositories = append(repositories, ossexplorer.Repository{
			Name:         r.GetName(),
			Forks:        r.GetForksCount(),
			Stars:        r.GetStargazersCount(),
			PullRequests: pulls,
		})
	}

	return repositories, nil
}

func (g *GitHubExplorer) getPullRequestCount(ctx context.Context, org, repo string) (int, error) {
	options := &gh.PullRequestListOptions{
		State: "all",
		ListOptions: gh.ListOptions{
			PerPage: 10,
		},
	}

	var numPulls int
	var done bool

	for !done {
		pulls, resp, err := g.client.PullRequests.List(ctx, org, repo, options)

		if err != nil {
			return -1, err
		}

		numPulls += len(pulls)

		done = resp.NextPage == 0
		options.Page = resp.NextPage
	}

	return numPulls, nil
}

func NewExplorer(accessToken string) *GitHubExplorer {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tokenClient := oauth2.NewClient(ctx, tokenSource)

	return &GitHubExplorer{client: gh.NewClient(tokenClient)}
}
