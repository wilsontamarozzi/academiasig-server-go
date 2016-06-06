package models

import(
	"github.com/satori/go.uuid"
)

type Usuario struct {
	UUID			uuid.UUID	`gorm:"primary_key;"`
	PessoaUUID 		uuid.UUID
	Ativo			bool
	Login			string
	Senha 			string
}

type Usuarios []Usuario