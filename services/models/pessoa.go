package models

import (
	_"time"
)

type Pessoa struct {
	Id 					int64				`gorm:"primary_key; AUTO_INCREMENT"`
	AdmCadastroId		int64
	//DataCadastro		time.Time
	Ativo				bool
	Email				string 
	Nome 				string
	Observacao			string 
	TipoPessoa			string
	Suporte				string 
	Bairro				string
	Cep					string 
	Complemento			string
	Logradouro			string 
	Numero				string 
	CidadeId			int64
	Sexo				bool
	//DataNascimento		time.Time
	Cpf 				string
	Rg					string
	TelefoneResidencial	string
	TelefoneCelular		string
	TelefoneEmpresa		string
	UsuarioSistema		bool
	EmpresaSistema		bool
	RazaoSocial			string
	Cnpj 				string
	TelefoneComercial	string
	Fax					string
	Website				string
	InscricaoEstadual	string
	InscricaoMunicipal	string
	Usuario				*Usuario 		`gorm:"ForeignKey:pessoa_id;AssociationForeignKey:id"`
}

type Pessoas []Pessoa