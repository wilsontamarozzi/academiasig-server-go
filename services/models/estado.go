package models

type Estado struct {
	EstadoId		int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	PaisId 			int64
	Nome			string	
	Sigla			string 	
	Pais 			Pais 	`gorm:"ForeignKey:pais_id;associationforeignkey:pais_id"`
}

type Estados []Estado