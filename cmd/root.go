package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	homeEnvVar = "VEGA_HOME"
)

var (
	flagDebug bool
	vegaHome  string
	rootCmd   *cobra.Command
)

func init() {
	rootCmd = newRootCmd(os.Stdout, os.Stdin)
}

func newRootCmd(out io.Writer, in io.Reader) *cobra.Command {
	const rootDesc = "vega - speed development"

	cmd := &cobra.Command{
		Use:   "vega",
		Short: rootDesc,
		Long:  rootDesc,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			os.Setenv(homeEnvVar, vegaHome)
			return
		},
	}

	persistentFlags := cmd.PersistentFlags()
	persistentFlags.StringVar(&vegaHome, "home", defaultVegaHome(), "location of your Vega init directory")
	persistentFlags.BoolVar(&flagDebug, "debug", false, "enable verbose output")

	cmd.AddCommand(newInitCmd(out, in))
	cmd.AddCommand(newHomeCmd(out))
	cmd.AddCommand(newCreateCmd(out))
	cmd.AddCommand(newVersionCmd(out))
	cmd.AddCommand(newStarterKitCmd(out))
	cmd.AddCommand(newUpCmd(out))
	cmd.AddCommand(newDownCmd(out))
	cmd.AddCommand(newRepoCmd(out))

	return cmd
}

func defaultVegaHome() string {
	if home := os.Getenv(homeEnvVar); home != "" {
		return home
	}

	homeEnvPath := os.Getenv("HOME")
	if homeEnvPath == "" && runtime.GOOS == "windows" {
		homeEnvPath = os.Getenv("USERPROFILE")
	}

	return filepath.Join(homeEnvPath, ".vega")
}

func homePath() string {
	return os.ExpandEnv(vegaHome)
}

func debug(format string, args ...interface{}) {
	if flagDebug {
		format = fmt.Sprintf("[debug] %s\n", format)
		fmt.Printf(format, args...)
	}
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
