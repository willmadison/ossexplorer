package mocks

import (
	"context"

	"github.com/willmadison/ossexplorer"
)

type StubExplorer struct {
	org   ossexplorer.Organization
	repos []ossexplorer.Repository
}

func (s StubExplorer) FindOrganization(context.Context, string) (ossexplorer.Organization, error) {
	return s.org, nil
}

func (s StubExplorer) FindRepositoriesFor(context.Context, ossexplorer.Organization, ...ossexplorer.RepositoryResultModifier) ([]ossexplorer.Repository, error) {
	return s.repos, nil
}

func NewStubExplorer(org ossexplorer.Organization, repos []ossexplorer.Repository) StubExplorer {
	return StubExplorer{org, repos}
}
