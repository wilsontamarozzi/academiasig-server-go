package services

import (
	"github.com/jinzhu/gorm"
	"AcademiaSIG-API/database"
)

var Con *gorm.DB

func init() {
	Con = database.GetConnection()
}