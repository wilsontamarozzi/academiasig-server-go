package models

type Banco struct {
	BancoId int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	Numero	string
	Nome	string
}

type Bancos []Banco