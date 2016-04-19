package models

type Usuario struct {
	Id 				uint64	`gorm:"primary_key; AUTO_INCREMENT"`
	PessoaId 		int64
	Ativo			bool
	Login			string
	Senha 			string
}

type Usuarios []Usuario