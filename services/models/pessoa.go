package models

import (
	_"time"
)

type Pessoa struct {
	Id 					int64				`gorm:"primary_key; AUTO_INCREMENT"`
	AdmCadastroId		int64
	LogradouroId 		int64
	//DataCadastro		time.Time
	Ativo				bool
	Email				string 
	Nome 				string
	Observacao			string 
	TipoPessoa			string
	Suporte				string 
	Complemento			string
	Numero				string
	Sexo				*bool
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
	Logradouro 			Logradouro 		`gorm:"ForeignKey:endereco_codigo;AssociationForeignKey:logradouro_id"`
}

type Pessoas []Pessoa

func (p Pessoa) IsValid() map[string][]string {

	err := make(map[string][]string)

	if p.Nome == "" {
		err["nome"] = append(err["nome"], "Nome não pode estar vázio.")
	}

	if len(p.Nome) < 2 {
		err["nome"] = append(err["nome"], "Nome deve ter mais de 2 caracteres.")
	}

	if p.TipoPessoa == "" {
		err["tipoPessoa"] = append(err["tipoPessoa"], "Tipo de Pessoa não definido.")
	} else {
		if p.TipoPessoa == "F" {
			if p.Sexo == nil {
				err["sexo"] = append(err["sexo"], "Sexo obrigatório.")
			}			
		}
	}

	return err
}