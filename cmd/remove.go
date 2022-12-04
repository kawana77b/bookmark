package cmd

import (
	"fmt"

	"github.com/kawana77b/bookmark/internal/controllers"

	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Removes the selected directory from the list of bookmark",
	Long:    ``,
	PreRunE: preRunCommon,
	RunE:    runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func runRemove(cmd *cobra.Command, args []string) error {
	c := &controllers.DirectoryController{}
	dir, err := chooseDirFromAllDirs(c)
	if err != nil {
		return err
	}

	if dir == nil || len(dir.Path) == 0 {
		return nil
	}

	fmt.Println("path: " + dir.Path)
	fmt.Println()

	fmt.Printf("Do you want to delete it? (y/n): ")

	var yn string
	_, err = fmt.Scan(&yn)
	if err != nil {
		return err
	}

	fmt.Println()

	if yn == "y" {
		err := c.Remove(dir)
		if err != nil {
			return err
		}

		fmt.Println(aurora.Yellow("remove from database"))
	}

	return nil
}
