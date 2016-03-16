package models

import (
	"time"
)

type Pessoa struct {
	PessoaId 		int64		`gorm:"column:pessoa_id; primary_key;" sql:"AUTO_INCREMENT" json:"pessoaId"`
	AdmCadastroId	int64 		`gorm:"column:adm_cadastro_id" json:"admCadastroId"`
	DataCadastro	time.Time 	`gorm:"column:data_cadastro" json:"dataCadastro"`
	Ativo			bool		`json:"ativo"`
	Email			string 		`json:"email"`
	Nome 			string		`json:"nome"`
	Observacao		string 		`json:"observacao"`
	TipoPessoa		string 		`json:"tipo_pessoa"`
	Suporte			string 		`json:"suporte"`
	Bairro			string 		`json:"bairro"`
	Cep				string 		`json:"cep"`
	Complemento		string 		`json:"complemento"`
	Logradouro		string 		`json:"logradouro"`
	Numero			string 		`json:"numero"`
	CidadeId		int64		`json:"cidade_id"`	
}

type Pessoas []Pessoa