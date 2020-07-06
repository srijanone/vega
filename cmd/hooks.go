package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newHooksCmd(out io.Writer) *cobra.Command {
	const repoDesc = "manage git hooks"

	repoCmd := &cobra.Command{
		Use:   "hooks",
		Short: repoDesc,
		Long:  repoDesc,
	}

	repoCmd.AddCommand(newHooksInstallCmd(out))

	return repoCmd
}
