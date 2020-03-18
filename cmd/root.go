package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	homeEnvVar = "VEGA_HOME"
)

var (
	flagDebug bool
	vegaHome  string
	rootCmd   *cobra.Command
	// TODO: globalConfig is the configuration stored in $VEGA_HOME/config.toml
	// globalConfig VegaConfig
)

func init() {
	fmt.Println("init of cmd")
	rootCmd = newRootCmd(os.Stdout, os.Stdin)
}

func newRootCmd(out io.Writer, in io.Reader) *cobra.Command {
	rootDesc := "vega - speed development"

	cmd := &cobra.Command{
		Use:   "vega",
		Short: rootDesc,
		Long:  rootDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command called")
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if flagDebug {
				log.SetLevel(log.DebugLevel)
			}
			// fmt.Printf("homeEnvVar", homeEnvVar)
			// fmt.Printf("vegaHome", vegaHome)
			os.Setenv(homeEnvVar, vegaHome)
			// globalConfig, err = ReadConfig()
			return
		},
	}

	cmd.AddCommand(newInitCmd(out))
	cmd.AddCommand(newHomeCmd(out))

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

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

func validateArgs(args, expectedArgs []string) error {
	if len(args) != len(expectedArgs) {
		return fmt.Errorf("This command needs %v argument(s): %v", len(expectedArgs), expectedArgs)
	}
	return nil
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
