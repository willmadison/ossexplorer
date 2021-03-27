package ossexplorer

import (
	"context"
	"sort"
)

type Explorer interface {
	FindOrganization(context.Context, string) (Organization, error)
	FindRepositoriesFor(context.Context, Organization, ...RepositoryResultModifier) ([]Repository, error)
}

type Organization struct {
	Name string
}

type Repository struct {
	Name                       string
	Stars, Forks, PullRequests int
}

func (r Repository) ContributionRate() float64 {
	return float64(r.PullRequests) / float64(r.Forks)
}

type RepositoryResultModifier func([]Repository) []Repository

func ByStarsDescending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Stars > repos[j].Stars
	})

	return repos
}

func ByStarsAscending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Stars < repos[j].Stars
	})

	return repos
}

func ByForksDescending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Forks > repos[j].Forks
	})

	return repos
}

func ByForksAscending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Forks < repos[j].Forks
	})

	return repos
}

func ByPullRequestsDescending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].PullRequests > repos[j].PullRequests
	})

	return repos
}

func ByPullRequestsAscending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].PullRequests < repos[j].PullRequests
	})

	return repos
}

func ByContributionRateDescending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].ContributionRate() > repos[j].ContributionRate()
	})

	return repos
}

func ByContributionRateAscending(repos []Repository) []Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].ContributionRate() < repos[j].ContributionRate()
	})

	return repos
}
