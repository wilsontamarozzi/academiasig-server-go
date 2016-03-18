package models

type Estado struct {
	EstadoId		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	PaisId 			int64
	EstadoNome		string	
	Sigla			string 	
	Pais 			Pais 	`gorm:"foreignkey:pais_id;associationforeignkey:pais_id"`
}

type Estados []Estado