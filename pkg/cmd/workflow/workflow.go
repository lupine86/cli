package workflow

import (
	cmdList "github.com/cli/cli/pkg/cmd/workflow/list"
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdWorkflow(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workflow <command>",
		Short: "Manage Actions workflows",
		Long:  "Work with Actions workflows",
	}

	cmd.AddCommand(cmdList.NewCmdList(f, nil))

	return cmd
}
