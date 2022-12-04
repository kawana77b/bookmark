package cmd

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/kawana77b/bookmark/internal/config"
	"github.com/kawana77b/bookmark/internal/database"
	"github.com/kawana77b/bookmark/internal/util"

	"github.com/spf13/cobra"
)

const (
	bash = "bash"
	fish = "fish"
	pwsh = "powershell"
)

var TplFs embed.FS

var shells []string = []string{bash, fish, pwsh}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Outputs a script according to the argument shell and makes this application usable",
	Long:    ``,
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), shellMatchRule),
	PreRunE: preRunInit,
	RunE:    runInit,
}

// シェルのプロファイル設定を出力するかのフラグ
var shouldDisplayShellProfile bool

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolVarP(&shouldDisplayShellProfile, "profile", "p", false, "output shell profile settings")
}

func shellMatchRule(cmd *cobra.Command, args []string) error {
	arg := args[0]
	for _, v := range shells {
		if arg == v {
			return nil
		}
	}

	return fmt.Errorf("argument must be one of %s", strings.Join(shells, ", "))
}

func preRunInit(cmd *cobra.Command, args []string) error {
	if !util.FileExists(config.DbPath()) {
		database.Initialize()
	}

	return nil
}

func runInit(cmd *cobra.Command, args []string) error {
	shell := args[0]

	var err error
	if shouldDisplayShellProfile {
		err = displayShellProfile(shell)
	} else {
		err = displayFunctions(shell)
	}

	if err != nil {
		return err
	}

	return nil
}

func displayShellProfile(shell string) error {
	tpl, err := template.ParseFS(TplFs, "template/profile.tpl")
	if err != nil {
		return err
	}

	var tplName string
	switch shell {
	case bash:
		tplName = "bash-prof"
	case fish:
		tplName = "fish-prof"
	case pwsh:
		tplName = "powershell-prof"
	}

	var b bytes.Buffer
	err = tpl.ExecuteTemplate(&b, tplName, nil)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", b.String())

	return nil
}

func displayFunctions(shell string) error {
	tpl, err := getTemplate(shell)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println(tpl)
	fmt.Println()

	return nil
}

func getTemplate(shell string) (string, error) {
	var data []byte
	var err error
	switch shell {
	case bash:
		data, err = TplFs.ReadFile("template/bash.sh")
	case fish:
		data, err = TplFs.ReadFile("template/fish.fish")
	case pwsh:
		data, err = TplFs.ReadFile("template/powershell.ps1")
	default:
		data, err = []byte{}, errors.New("template not found")
	}

	return string(data), err
}
