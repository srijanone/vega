package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newRepoCmd(out io.Writer) *cobra.Command {
	const repoDesc = "manage repositories of starterkit"

	repoCmd := &cobra.Command{
		Use:   "repo",
		Short: repoDesc,
		Long:  repoDesc,
	}

	repoCmd.AddCommand(newAddCmd(out))
	repoCmd.AddCommand(newRepositoryCmdList(out))

	return repoCmd
}
