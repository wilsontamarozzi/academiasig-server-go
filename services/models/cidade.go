package models

type Cidade struct {
	CidadeId 		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT" json:"cidadeId"`
	EstadoId		int64	`json:"estadoId"`
	PaisId			int64	`json:"paisId"`
	Nome			string 	`gorm:"column:cidade_nome" json:"nome"`
	Uf				string	`json:"uf"`
	Pais 			string	`json:"pais"`
	CodigoSisbacen	int64	`json:"codigoSisbacen"`
	CodigoIbge		int64	`json:"codigoIbge"`
	CodigoSiafi		int64	`json:"codigoSiafi"`
	CodigoSetec		int64	`json:"codigoSetec"`
	Estado 			Estado 	`gorm:"foreignkey:estado_id;associationforeignkey:estado_id" json:"estado"`
}

type Cidades []Cidade