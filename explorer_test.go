package ossexplorer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/ossexplorer"
)

func ItProperlySortsFoundRepositoriesByStarsDescending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	sorted := ossexplorer.ByStarsDescending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}

func ItProperlySortsFoundRepositoriesByStarsAscending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	sorted := ossexplorer.ByStarsAscending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}

func ItProperlySortsFoundRepositoriesByForksDescending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	sorted := ossexplorer.ByForksDescending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}

func ItProperlySortsFoundRepositoriesByForksAscending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	sorted := ossexplorer.ByForksAscending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}

func ItProperlySortsFoundRepositoriesByPullRequestsDescending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	sorted := ossexplorer.ByPullRequestsDescending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}

func ItProperlySortsFoundRepositoriesByPullRequestsAscending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	sorted := ossexplorer.ByPullRequestsAscending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}

func ItProperlySortsFoundRepositoriesByContributionRateDescending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	sorted := ossexplorer.ByContributionRateDescending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}

func ItProperlySortsFoundRepositoriesByContributionRateAscending(t *testing.T) {
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	sorted := ossexplorer.ByContributionRateAscending(repos)

	expected := []ossexplorer.Repository{
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
	}

	assert.ElementsMatch(t, expected, sorted, "we should have sorted the respositories properly")
}
