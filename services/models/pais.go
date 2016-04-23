package models

type Pais struct {
	Iso		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	Iso3	string
	Numcode	string
	Nome	string
}

type Paises []Pais