package models

import (
	"time"
)

type Pessoa struct {
	PessoaId 		int64				`gorm:"primary_key; AUTO_INCREMENT" json:"pessoaId"`
	AdmCadastroId	int64 				`json:"admCadastroId"`
	DataCadastro	time.Time 			`json:"dataCadastro"`
	Ativo			bool				`json:"ativo"`
	Email			string 				`json:"email"`
	Nome 			string				`json:"nome"`
	Observacao		string 				`json:"observacao"`
	TipoPessoa		string 				`json:"tipoPessoa"`
	Suporte			string 				`json:"suporte"`
	Bairro			string 				`json:"bairro"`
	Cep				string 				`json:"cep"`
	Complemento		string 				`json:"complemento"`
	Logradouro		string 				`json:"logradouro"`
	Numero			string 				`json:"numero"`
	CidadeId		int64				`json:"cidadeId"`
	PessoaFisica	*PessoaFisica 		`gorm:"foreignkey:pessoa_id; associationforeignkey:pessoa_id"`
	PessoaJuridica	*PessoaJuridica		`gorm:"foreignkey:pessoa_id; associationforeignkey:pessoa_id"`
}

type Pessoas []Pessoa