package models

import (
	"time"
)

type PessoaFisica struct {
	PessoaFisicaId 		uint64 		`gorm:"primary_key; AUTO_INCREMENT"`
	PessoaId 			int64
	Sexo				bool
	DataNascimento		time.Time
	Cpf 				string
	Rg					string
	TelefoneResidencial	string
	TelefoneCelular		string
	TelefoneEmpresa		string
	UsuarioSistema		bool
	Usuario				*Usuario 	`gorm:"ForeignKey:pessoa_fisica_id;AssociationForeignKey:pessoa_fisica_id"`
}

type PessoasFisica []PessoaFisica