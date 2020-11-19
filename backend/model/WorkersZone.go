package model

import "github.com/jinzhu/gorm"

type WorkersZone struct { // Lista de workers a trabalhar em cada zona
	gorm.Model `swaggerignore:"true"`
	IdWorker   int "gorm:\"foreignKey:Worker_ID\""
	IdZone     int "gorm:\"foreignKey:Zone_ID\""
}
