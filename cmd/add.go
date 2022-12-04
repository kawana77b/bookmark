package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kawana77b/bookmark/internal/controllers"
	"github.com/kawana77b/bookmark/internal/models"

	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Bookmark the current working directory",
	Long:    ``,
	PreRunE: preRunCommon,
	RunE:    runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	c := &controllers.DirectoryController{}

	path, _ := filepath.Abs(wd)
	if d, _ := c.FindByPath(path); d != nil {
		return errors.New("directory is already registered")
	}

	d := models.NewDirecotory(path)
	err = c.Add(d)
	if err != nil {
		return err
	}

	fmt.Printf("bookmarked: %s\n", aurora.Yellow(wd))

	return nil
}
