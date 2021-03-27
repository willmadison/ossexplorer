package cli_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/ossexplorer"
	"github.com/willmadison/ossexplorer/cli"
	"github.com/willmadison/ossexplorer/mocks"
)

func TestItReturnsSuccessForProperlyFormattedArguments(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg", "stars", "asc", "1"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	env := cli.Environment{
		Stderr: stderr,
		Stdout: stdout,
		Stdin:  stdin,
	}

	quit := make(chan struct{})

	org := ossexplorer.Organization{Name: "DummyOrg"}
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	returnCode := 0

	go func() {
		returnCode = cli.Run(mocks.NewStubExplorer(org, repos), env)
		close(quit)
	}()

	time.Sleep(30 * time.Millisecond)

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}

func TestItDisplaysFoundReposSortedAppropriately(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	env := cli.Environment{
		Stderr: stderr,
		Stdout: stdout,
		Stdin:  stdin,
	}

	quit := make(chan struct{})

	org := ossexplorer.Organization{Name: "DummyOrg"}
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	returnCode := 0

	go func() {
		returnCode = cli.Run(mocks.NewStubExplorer(org, repos), env)
		close(quit)
	}()

	time.Sleep(30 * time.Millisecond)

	expectedOutput := ` # Repository # Stars # Forks # Pull Requests Contribution %
 _ __________ _______ _______ _______________ ______________
 1 DummyRepo2      10      25               2          8.00%
 2 DummyRepo1       5       6              10        166.67%
`
	assert.Equal(t, expectedOutput, stdout.String(), "we should have gotten the appropriate output")

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}

func TestItDisplaysLimitedResultsAccordingToGivenConstraints(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg", "stars", "desc", "1"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	env := cli.Environment{
		Stderr: stderr,
		Stdout: stdout,
		Stdin:  stdin,
	}

	quit := make(chan struct{})

	org := ossexplorer.Organization{Name: "DummyOrg"}
	repos := []ossexplorer.Repository{
		{Name: "DummyRepo1", Stars: 5, Forks: 6, PullRequests: 10},
		{Name: "DummyRepo2", Stars: 10, Forks: 25, PullRequests: 2},
	}

	returnCode := 0

	go func() {
		returnCode = cli.Run(mocks.NewStubExplorer(org, repos), env)
		close(quit)
	}()

	time.Sleep(30 * time.Millisecond)

	expectedOutput := ` # Repository # Stars # Forks # Pull Requests Contribution %
 _ __________ _______ _______ _______________ ______________
 1 DummyRepo2      10      25               2          8.00%
`
	assert.Equal(t, expectedOutput, stdout.String(), "we should have gotten the appropriate output")

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}

func TestItDisplaysAnUnableToFindOrganizationMessageInTheEventThatWeCannotFindAnOrg(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg", "stars", "desc", "1"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	env := cli.Environment{
		Stderr: stderr,
		Stdout: stdout,
		Stdin:  stdin,
	}

	quit := make(chan struct{})

	returnCode := 0

	go func() {
		returnCode = cli.Run(mocks.NewFailAlwaysExplorer(), env)
		close(quit)
	}()

	time.Sleep(30 * time.Millisecond)

	expectedOutput := `Unable to locate the given organization (DummyOrg). Please ensure you have access and that it exists.
`

	assert.Equal(t, expectedOutput, stderr.String(), "we should have gotten the appropriate output")

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}

func TestItDisplaysARepositoryLookupFailureMessageInTheEventThatWeCannotFindAnyRepositoriesForAGivenOrg(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg", "stars", "desc", "1"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	env := cli.Environment{
		Stderr: stderr,
		Stdout: stdout,
		Stdin:  stdin,
	}

	quit := make(chan struct{})

	returnCode := 0

	go func() {
		returnCode = cli.Run(mocks.NewErroneousExplorer(nil, errors.New("repo lookup failure")), env)
		close(quit)
	}()

	time.Sleep(30 * time.Millisecond)

	expectedOutput := `Repository lookup failed for DummyOrg. Please ensure you have read access to this organization's repositories.
`
	assert.Equal(t, expectedOutput, stderr.String(), "we should have gotten the appropriate output")

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}

func TestItDisplaysANoRepositoriesFoundMessageWhenAnOrgHasNoRepos(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg", "stars", "desc", "1"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	env := cli.Environment{
		Stderr: stderr,
		Stdout: stdout,
		Stdin:  stdin,
	}

	quit := make(chan struct{})

	org := ossexplorer.Organization{Name: "DummyOrg"}
	repos := []ossexplorer.Repository{}

	returnCode := 0

	go func() {
		returnCode = cli.Run(mocks.NewStubExplorer(org, repos), env)
		close(quit)
	}()

	time.Sleep(30 * time.Millisecond)

	expectedOutput := `No repositories found for DummyOrg. Nothing to do here but chill 😎...
`
	assert.Equal(t, expectedOutput, stdout.String(), "we should have gotten the appropriate output")

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}

func TestItDisplaysANoRepositoriesFoundMessageWhenAnOrgHasNilRepos(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg", "stars", "desc", "1"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	env := cli.Environment{
		Stderr: stderr,
		Stdout: stdout,
		Stdin:  stdin,
	}

	quit := make(chan struct{})

	org := ossexplorer.Organization{Name: "DummyOrg"}
	var repos []ossexplorer.Repository

	returnCode := 0

	go func() {
		returnCode = cli.Run(mocks.NewStubExplorer(org, repos), env)
		close(quit)
	}()

	time.Sleep(30 * time.Millisecond)

	expectedOutput := `No repositories found for DummyOrg. Nothing to do here but chill 😎...
`
	assert.Equal(t, expectedOutput, stdout.String(), "we should have gotten the appropriate output")

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}
