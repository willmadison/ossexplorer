package mocks

import (
	"context"
	"errors"

	"github.com/willmadison/ossexplorer"
)

type StubExplorer struct {
	org   ossexplorer.Organization
	repos []ossexplorer.Repository
}

func (s StubExplorer) FindOrganization(context.Context, string) (ossexplorer.Organization, error) {
	return s.org, nil
}

func (s StubExplorer) FindRepositoriesFor(ctx context.Context, org ossexplorer.Organization, modifiers ...ossexplorer.RepositoryResultModifier) ([]ossexplorer.Repository, error) {
	repositories := s.repos

	if len(modifiers) > 0 {
		for _, mod := range modifiers {
			if mod != nil {
				repositories = mod(repositories)
			}
		}
	}

	return repositories, nil
}

func NewStubExplorer(org ossexplorer.Organization, repos []ossexplorer.Repository) StubExplorer {
	return StubExplorer{org, repos}
}

type FailAlwaysExplorer struct {
}

func (f FailAlwaysExplorer) FindOrganization(context.Context, string) (ossexplorer.Organization, error) {
	return ossexplorer.Organization{}, errors.New("always fails")
}

func (f FailAlwaysExplorer) FindRepositoriesFor(ctx context.Context, org ossexplorer.Organization, modifiers ...ossexplorer.RepositoryResultModifier) ([]ossexplorer.Repository, error) {
	return nil, errors.New("always fails")
}

func NewFailAlwaysExplorer() FailAlwaysExplorer {
	return FailAlwaysExplorer{}
}

type ErroneousExplorer struct {
	orgError, repoError error
}

func (e ErroneousExplorer) FindOrganization(context.Context, string) (ossexplorer.Organization, error) {
	return ossexplorer.Organization{}, e.orgError
}

func (e ErroneousExplorer) FindRepositoriesFor(ctx context.Context, org ossexplorer.Organization, modifiers ...ossexplorer.RepositoryResultModifier) ([]ossexplorer.Repository, error) {
	return nil, e.repoError
}

func NewErroneousExplorer(orgError, repoError error) ErroneousExplorer {
	return ErroneousExplorer{orgError, repoError}
}
