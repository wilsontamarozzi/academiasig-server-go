package models

type Cidade struct {
	CidadeId 		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	EstadoId		int64	
	PaisId			int64	
	Nome			string 	`gorm:"column:cidade_nome"`
	Uf				string	
	Pais 			string	
	CodigoSisbacen	int64	
	CodigoIbge		int64	
	CodigoSiafi		int64	
	CodigoSetec		int64	
	Estado 			Estado 	`gorm:"foreignkey:estado_id;associationforeignkey:estado_id"`
}

type Cidades []Cidade