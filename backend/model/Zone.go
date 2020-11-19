package model

import "github.com/jinzhu/gorm"

type Zone struct {
	gorm.Model `swaggerignore:"true"`
	Zone_ID    int    "gorm:primaryKey"
	Name       string "gorm:unique"
	Latitude   float32
	Longitude  float32
	Limit      int
}
