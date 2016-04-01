package services

import (
	"github.com/jinzhu/gorm"
	"academiasig-api/database"
)

var Con *gorm.DB

func init() {
	Con = database.GetConnection()
}