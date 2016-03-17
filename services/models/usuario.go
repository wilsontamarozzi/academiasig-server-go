package models

type Usuario struct {
	UsuarioId 		int64		`gorm:"primary_key; AUTO_INCREMENT" json:"usuarioId"`
	PessoaFisicaId 	int64 		`json:"pessoaFisicaId"`
	Ativo			bool		`json:"ativo"`
	Login			string		`json:"login"`
	Senha 			string		`json:"senha"`
}

type Usuarios []Usuario