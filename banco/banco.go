package banco

type Banco struct {
	BancoId int64	`gorm:"column:banco_id; primary_key;" sql:"AUTO_INCREMENT" json:"bancoId"`
	Numero	string 	`json:"numero"`
	Nome	string 	`json:"nome"`
}

type Bancos []Banco