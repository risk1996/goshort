package models

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Path        string         `json:"path" gorm:"uniqueIndex"`
	Target      string         `json:"target"`
	AdminSecret string         `json:"admin_secret" gorm:"uniqueIndex"`
}
