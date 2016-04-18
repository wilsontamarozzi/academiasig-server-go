package models

import (
	"time"
)

type Pessoa struct {
	Id 				int64				`gorm:"primary_key; AUTO_INCREMENT"`
	AdmCadastroId	int64
	DataCadastro	time.Time
	Ativo			bool
	Email			string 
	Nome 			string
	Observacao		string 
	TipoPessoa		string
	Suporte			string 
	Bairro			string
	Cep				string 
	Complemento		string
	Logradouro		string 
	Numero			string 
	CidadeId		int64
	PessoaFisica	*PessoaFisica		`gorm:"ForeignKey:pessoa_id;AssociationForeignKey:id"`
	PessoaJuridica	*PessoaJuridica		`gorm:"ForeignKey:pessoa_id;AssociationForeignKey:id"`
}

type Pessoas []Pessoa