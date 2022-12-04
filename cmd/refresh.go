package cmd

import (
	"fmt"

	"github.com/kawana77b/bookmark/internal/controllers"
	"github.com/kawana77b/bookmark/internal/models"

	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:     "refresh",
	Short:   "Removes non-existent directories from the database",
	Long:    ``,
	PreRunE: preRunCommon,
	RunE:    runRefresh,
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}

func runRefresh(cmd *cobra.Command, args []string) error {
	c := &controllers.DirectoryController{}
	dirs, err := c.FindAll()
	if err != nil {
		return err
	}

	rmDirs := []*models.Directory{}
	for _, v := range dirs {
		if !v.Exists() {
			rmDirs = append(rmDirs, v)
		}
	}

	err = c.RemoveRange(rmDirs)
	if err != nil {
		return err
	}

	fmt.Println(aurora.Cyan("database has been refreshed"))

	return nil
}
