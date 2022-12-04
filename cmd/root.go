package cmd

import (
	"fmt"
	"os"

	"github.com/kawana77b/bookmark/internal/config"

	"github.com/spf13/cobra"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   config.APP_NAME,
	Short: fmt.Sprintf(`%s is a tool that allows you to bookmark directories on the CLI command line`, config.APP_NAME),
	Long: fmt.Sprintf(`%s is a tool that allows you to bookmark directories on the CLI command line.

Functions are provided by subcommands.
	`, config.APP_NAME),
}

func Execute() {
	rootCmd.Version = Version

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
