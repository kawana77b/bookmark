package cmd

import (
	"fmt"

	"github.com/kawana77b/bookmark/internal/config"
	"github.com/kawana77b/bookmark/internal/proc"
	"github.com/kawana77b/bookmark/internal/util"

	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the database path and other information",
	Long:  ``,
	Run:   runStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func runStatus(cmd *cobra.Command, args []string) {
	fmt.Printf("DATABASE PASS:\n")
	fmt.Printf("  %s\n", config.DbPath())

	fmt.Println()

	fmt.Printf("DATABASE EXISTS:\n")
	fmt.Printf("  %s\n", dispBool(util.FileExists(config.DbPath())))

	fmt.Println()

	canUsePeco := func() bool {
		if err := proc.CheckCanUsePeco(); err != nil {
			return false
		} else {
			return true
		}
	}()

	fmt.Printf("PECO INSTALLED:\n")
	fmt.Printf("  %s\n", dispBool(canUsePeco))
}

func dispBool(b bool) string {
	var v aurora.Value
	if b {
		v = aurora.BrightGreen("%v")
	} else {
		v = aurora.BrightRed("%v")
	}
	return aurora.Sprintf(v, b)
}
