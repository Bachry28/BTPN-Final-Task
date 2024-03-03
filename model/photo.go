package model

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `gorm:"foreignKey:UserID;references:ID"`
}
