package models

type Estado struct {
	EstadoId		int64	`gorm:"column:estado_id; primary_key;" sql:"AUTO_INCREMENT" json:"cidadeId"`
	PaisId 			int64	`json:"paisId"`
	EstadoNome		string	`json:"nome"`
	Sigla			string 	`json:"sigla"`
}

type Estados []Estado