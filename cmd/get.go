package cmd

import (
	"fmt"

	"github.com/kawana77b/bookmark/internal/controllers"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Find bookmarked directories",
	Long:    ``,
	PreRunE: preRunCommon,
	RunE:    runGet,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func runGet(cmd *cobra.Command, args []string) error {
	c := &controllers.DirectoryController{}
	dir, err := chooseDirFromAllDirs(c)
	if err != nil {
		return err
	}

	if dir != nil && dir.Exists() {
		fmt.Printf("%s\n", dir.Path)
	} else {
		fmt.Println("")
	}

	return nil
}
