package models

type Pais struct {
	Id		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	Iso		string	
	Iso3	string
	Numcode	string
	Nome	string
}

type Paises []Pais