package cli_test

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/ossexplorer"
	"github.com/willmadison/ossexplorer/cli"
	"github.com/willmadison/ossexplorer/mocks"
)

type writeNotifier struct {
	*sync.Cond
}

func (w *writeNotifier) Write(p []byte) (n int, err error) {
	defer func() {
		time.Sleep(10 * time.Millisecond)
		w.Broadcast()
	}()
	return 0, nil
}

func TestItReturnsSuccessForProperlyFormattedArguments(t *testing.T) {
	os.Args = []string{"explore", "repos", "DummyOrg", "stars", "asc"}

	stdin, _ := io.Pipe()
	stdout := &bytes.Buffer{}

	outMutex := &sync.Mutex{}
	outnotify := &writeNotifier{sync.NewCond(outMutex)}

	stderr := &bytes.Buffer{}
	errMutex := &sync.Mutex{}
	errnotify := &writeNotifier{sync.NewCond(errMutex)}

	env := cli.Environment{
		Stderr: io.MultiWriter(stderr, errnotify),
		Stdout: io.MultiWriter(stdout, outnotify),
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

	<-quit
	assert.Equal(t, 0, returnCode, "there should be no errors here")
}
