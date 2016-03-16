package models

type Cidade struct {
	CidadeId 		int64	`gorm:"column:cidade_id; primary_key;" sql:"AUTO_INCREMENT" json:"cidadeId"`
	EstadoId		int64	`gorm:"column:estado_id" json:"estadoId"`
	PaisId			int64	`gorm:"column:pais_id" json:"paisId"`
	Nome			string 	`gorm:"column:cidade_nome" json:"nome"`
	Uf				string	`json:"uf"`
	Pais 			string	`json:"pais"`
	CodigoSisbacen	int64	`json:"codigoSisbacen"`
	CodigoIbge		int64	`json:"codigoIbge"`
	CodigoSiafi		int64	`json:"codigoSiafi"`
	CodigoSetec		int64	`json:"codigoSetec"`
}

type Cidades []Cidade