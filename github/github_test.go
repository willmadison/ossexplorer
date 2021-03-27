package github_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/ossexplorer"
	"github.com/willmadison/ossexplorer/github"
)

func TestItLocatesOrganizations(t *testing.T) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	if accessToken == "" {
		t.Skip("no GITHUB_ACCESS_TOKEN configured in this environment, skipping this test")
	}

	explorer := github.NewExplorer(accessToken)

	org, err := explorer.FindOrganization(context.Background(), "BlacksInTechnologyOrg")

	assert.Nil(t, err)
	assert.Equal(t, "BlacksInTechnologyOrg", org.Name)
}

func TestItFindsRepositoriesByOrganization(t *testing.T) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	if accessToken == "" {
		t.Skip("no GITHUB_ACCESS_TOKEN configured in this environment, skipping this test")
	}

	explorer := github.NewExplorer(accessToken)

	repos, err := explorer.FindRepositoriesFor(context.Background(), ossexplorer.Organization{Name: "BlacksInTechnologyOrg"})

	assert.Nil(t, err)
	assert.NotEmpty(t, repos, "we should have found at least one repository")
	assert.Contains(t, repos, ossexplorer.Repository{Name: "bit-slack-greeting-bot", Stars: 4, Forks: 4, PullRequests: 13})
}
