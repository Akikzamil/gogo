package model

import "gorm.io/gorm"

type Migration struct {
	gorm.Model
	Migration string `gorm:"migration"`
	Batch int `gorm:"batch"`
}