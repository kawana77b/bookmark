package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kawana77b/bookmark/internal/config"
	"github.com/kawana77b/bookmark/internal/models"
	"github.com/kawana77b/bookmark/internal/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// データベースと接続し、grom.DBのポインタを取得します
func Conn() (*gorm.DB, error) {
	return conn(config.DbPath())
}

func conn(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		CreateBatchSize: 100,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// データベースが作成された時に発火します
var OnDatabaseCreated func(dbPath string)

// データベースがマイグレートされた時に発火します
var OnDatabaseMigrated func()

// データベースを初期化します
func Initialize() error {
	if util.FileExists(config.DbPath()) {
		return fmt.Errorf("database file already exists: %s", config.DbPath())
	}

	if d := filepath.Dir(config.DbPath()); !util.DirExists(d) {
		os.MkdirAll(d, 0776)
	}

	db, err := Conn()
	if err != nil {
		return err
	}

	if OnDatabaseCreated != nil {
		OnDatabaseCreated(config.DbPath())
	}

	models := []interface{}{
		&models.Directory{},
	}

	for _, m := range models {
		db.AutoMigrate(m)
	}

	if OnDatabaseMigrated != nil {
		OnDatabaseMigrated()
	}

	return nil
}
