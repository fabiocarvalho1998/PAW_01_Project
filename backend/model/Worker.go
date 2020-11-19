package model

import "github.com/jinzhu/gorm"

type Worker struct {
	gorm.Model `swaggerignore:"true"`
	Worker_ID  int    "gorm:primaryKey"
	Username   string "gorm:unique"
	Password   string
	Name       string "gorm:unique"
	Admin      bool
}
