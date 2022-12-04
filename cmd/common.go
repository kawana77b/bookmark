package cmd

import (
	"errors"
	"strings"

	"github.com/kawana77b/bookmark/internal/config"
	"github.com/kawana77b/bookmark/internal/controllers"
	"github.com/kawana77b/bookmark/internal/models"
	"github.com/kawana77b/bookmark/internal/proc"
	"github.com/kawana77b/bookmark/internal/util"

	"github.com/spf13/cobra"
)

// DBパスとPECOインストールをチェックし、条件が満たされていなければエラー
func preRunCommon(cmd *cobra.Command, args []string) error {
	if !util.FileExists(config.DbPath()) {
		return errors.New("database file is not exists")
	}

	if err := proc.CheckCanUsePeco(); err != nil {
		return err
	}

	return nil
}

// 全てのディレクトリをデータベースから取得し、選択された1つを返す
func chooseDirFromAllDirs(c *controllers.DirectoryController) (*models.Directory, error) {
	dirs, err := c.FindAll()
	if err != nil {
		return nil, err
	}

	if len(dirs) == 0 {
		return nil, errors.New("bookmark is empty")
	}

	paths := []string{}
	for _, v := range dirs {
		paths = append(paths, v.Path)
	}

	data := strings.Join(paths, "\n")
	res := proc.ExecPeco([]byte(data))
	res = strings.TrimSpace(strings.ReplaceAll(res, "\n", ""))

	for _, v := range dirs {
		if res == v.Path {
			return v, nil
		}
	}

	return nil, nil
}
