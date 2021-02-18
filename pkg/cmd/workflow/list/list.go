package list

import (
	"net/http"

	"github.com/cli/cli/pkg/cmdutil"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/utils"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	IO         *iostreams.IOStreams
	HTTPClient func() (*http.Client, error)
}

func NewCmdList(f *cmdutil.Factory, runF func(*ListOptions) error) *cobra.Command {
	opts := &ListOptions{
		HTTPClient: f.HttpClient,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Action workflows in your GitHub account",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if runF != nil {
				return runF(opts)
			}
			return listRun(opts)
		},
	}

	return cmd
}

func listRun(opts *ListOptions) error {
	// GET /repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs

	//apiClient, err := opts.HTTPClient()
	//if err != nil {
	//	return err
	//}

	t := utils.NewTablePrinter(opts.IO)
	// cs := opts.IO.ColorScheme()
	// now := time.Now()

	t.AddField("Hello Field", nil, nil)

	t.EndRow()

	return t.Render()
}
