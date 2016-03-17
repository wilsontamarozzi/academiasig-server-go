package models

import (
	"time"
)

type PessoaFisica struct {
	PessoaFisicaId 		int64		`gorm:"primary_key; AUTO_INCREMENT" json:"pessoaFisicaId"`
	PessoaId 			int64 		`json:"pessoaId"`
	Sexo				bool		`json:"sexo"`
	DataNascimento		time.Time 	`json:"dataNascimento"`
	Cpf 				string		`json:"cpf"`
	Rg					string		`json:"rg"`
	TelefoneResidencial	string		`json:"telefoneResidencial"`
	TelefoneCelular		string		`json:"telefoneCelular"`
	TelefoneEmpresa		string		`json:"telefoneEmpresa"`
	UsuarioSistema		bool		`json:"usuarioSistema"`
	Usuario				*Usuario 	`gorm:"foreignkey:pessoa_fisica_id; associationforeignkey:pessoa_fisica_id"`
}

type PessoasFisica []PessoaFisica