package list

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/pkg/httpmock"
	"github.com/cli/cli/test"
	"github.com/stretchr/testify/assert"
)

func runCommand(rt http.RoundTripper, isTTY bool, cli string) (*test.CmdOut, error) {
	return nil, nil
}

func initFakeHTTP() *httpmock.Registry {
	return &httpmock.Registry{}
}

func TestPRList(t *testing.T) {
	http := initFakeHTTP()
	defer http.Verify(t)

	http.Register(httpmock.GraphQL(`query PullRequestList\b`), httpmock.FileResponse("./fixtures/prList.json"))

	output, err := runCommand(http, true, "")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, heredoc.Doc(`

		Showing 3 of 3 open pull requests in OWNER/REPO

		#32  New feature            feature
		#29  Fixed bad bug          hubot:bug-fix
		#28  Improve documentation  docs
	`), output.String())
	assert.Equal(t, ``, output.Stderr())
}

func TestListRun(t *testing.T) {
	tests := []struct {
		name       string
		opts       ListOptions
		isTTY      bool
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		{
			name: "list tty",
			opts: ListOptions{
				HTTPClient: func() (*http.Client, error) {
					createdAt := time.Now().Add(time.Duration(-24) * time.Hour)
					reg := &httpmock.Registry{}
					reg.Register(
						httpmock.REST("GET", "user/keys"),
						httpmock.StringResponse(fmt.Sprintf(`[
							{
								"id": 1234,
								"key": "ssh-rsa AAAABbBB123",
								"title": "Mac",
								"created_at": "%[1]s"
							},
							{
								"id": 5678,
								"key": "ssh-rsa EEEEEEEK247",
								"title": "hubot@Windows",
								"created_at": "%[1]s"
							}
						]`, createdAt.Format(time.RFC3339))),
					)
					return &http.Client{Transport: reg}, nil
				},
			},
			isTTY: true,
			wantStdout: heredoc.Doc(`
				Mac            ssh-rsa AAAABbBB123  1d
				hubot@Windows  ssh-rsa EEEEEEEK247  1d
			`),
			wantStderr: "",
		},
	}
}
