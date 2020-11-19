package model

import "github.com/jinzhu/gorm"

type Access struct {
	gorm.Model `swaggerignore:"true"`
	IdWorker   int "gorm:\"foreignKey:Worker_ID\""
	IdZone     int "gorm:\"foreignKey:Zone_ID\""
}
