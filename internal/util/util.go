package util

import (
	"os"
)

// ファイルが存在するか調べます
func FileExists(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
}

// ディレクトリが存在するか調べます
func DirExists(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}
