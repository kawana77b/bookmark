package controllers

import (
	"sort"

	"github.com/kawana77b/bookmark/internal/database"
	"github.com/kawana77b/bookmark/internal/models"
	"gorm.io/gorm"
)

/*----------------------------
* error
*----------------------------*/

type DirectoryNotFoundError struct {
	Path string
}

func (e DirectoryNotFoundError) Error() string {
	return "directory is not found: " + e.Path
}

/*----------------------------
* controller
*----------------------------*/

type DirectoryController struct {
}

func (d *DirectoryController) FindAll() ([]*models.Directory, error) {
	var dirs []*models.Directory = []*models.Directory{}

	db, err := database.Conn()
	if err != nil {
		return dirs, err
	}

	tx := db.Find(&dirs)
	if tx.Error != err {
		return dirs, err
	}

	sort.Slice(dirs, func(i int, j int) bool { return len(dirs[i].Path) < len(dirs[j].Path) })

	return dirs, nil
}

func (d *DirectoryController) FindByPath(path string) (*models.Directory, error) {
	db, err := database.Conn()
	if err != nil {
		return nil, err
	}

	var dir = &models.Directory{}
	tx := db.Limit(1).Find(&dir, "path = ?", path)
	if tx.RowsAffected > 0 {
		return dir, nil
	}

	return nil, DirectoryNotFoundError{Path: path}
}

func (d *DirectoryController) Add(dir *models.Directory) error {
	db, err := database.Conn()
	if err != nil {
		return err
	}

	tx := db.Create(dir)
	if tx.Error != nil {
		return err
	}

	return nil
}

func (d *DirectoryController) Remove(dir *models.Directory) error {
	db, err := database.Conn()
	if err != nil {
		return err
	}

	tx := db.Delete(&dir)
	if tx.Error != nil {
		return err
	}

	return nil
}

func (d *DirectoryController) RemoveRange(dirs []*models.Directory) error {
	db, err := database.Conn()
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		for _, v := range dirs {
			_tx := tx.Delete(v)

			if _tx.Error != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
