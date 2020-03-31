package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newRepoCmd(out io.Writer) *cobra.Command {
	const repoDesc = "Manage starterkits repositories"

	repoCmd := &cobra.Command{
		Use:   "repo",
		Short: repoDesc,
		Long:  repoDesc,
	}
	repoCmd.AddCommand(newAddCmd(out))
	repoCmd.AddCommand(newRepositoryList(out))
	return repoCmd
}
