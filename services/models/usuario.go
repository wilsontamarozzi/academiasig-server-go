package models

type Usuario struct {
	UsuarioId 		int64		`gorm:"primary_key; AUTO_INCREMENT"`
	PessoaFisicaId 	int64
	Ativo			bool
	Login			string
	Senha 			string
}

type Usuarios []Usuario