package models

import (
	_ "github.com/jinzhu/gorm"
)

type Role struct {
	Id          int
	Title       string
	Description string
	Status      int
	AddTime     int
}

func (r *Role) TableName() string {
	return "role"
}
