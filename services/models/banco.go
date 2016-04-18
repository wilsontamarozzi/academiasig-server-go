package models

type Banco struct {
	Id 		int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	Numero	string
	Nome	string
}

type Bancos []Banco