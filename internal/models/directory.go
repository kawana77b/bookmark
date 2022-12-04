package models

import (
	"os"
	"time"
)

type Directory struct {
	Id        uint `gorm:"primaryKey"`
	Path      string
	CreatedAt time.Time
}

func NewDirecotory(path string) *Directory {
	return &Directory{
		Path:      path,
		CreatedAt: time.Now(),
	}
}

func (d *Directory) Exists() bool {
	if fi, err := os.Stat(d.Path); err == nil && fi.IsDir() {
		return true
	}
	return false
}
