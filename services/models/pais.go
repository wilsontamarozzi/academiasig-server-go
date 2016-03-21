package models

type Pais struct {
	PaisId		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	Nome		string
}

type Paises []Pais