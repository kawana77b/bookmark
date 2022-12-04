package config

import (
	"os"
	"path/filepath"
)

const (
	APP_NAME = "bookmark"
	DB_NAME  = "data.db"
)

// データベース配置位置の固定絶対パスを返します
func DbPath() string {
	d, _ := os.UserCacheDir()
	return filepath.Join(d, APP_NAME, DB_NAME)
}
