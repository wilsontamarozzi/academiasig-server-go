package models

type Estado struct {
	UfCodigo		int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	UfSigla 		string
	UfDescricao		string
}

type Estados []Estado