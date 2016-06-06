package models

import (
	"time"
	"github.com/satori/go.uuid"
)

type Pessoa struct {
	UUID 				uuid.UUID		`gorm:"primary_key:true"`
	Codigo 				int64
	DataCadastro		time.Time 		`gorm:"default:'NOW()'"`
	Ativo				bool
	Email				string 
	Nome 				string
	Observacao			string 
	TipoPessoa			string
	Suporte				string 
	Complemento			string
	Numero				string
	Sexo				*bool
	DataNascimento		time.Time
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

	AdmCadastroId		int64
	LogradouroId 		int64

	Usuario				*Usuario 		`gorm:"ForeignKey:pessoa_uuid;AssociationForeignKey:uuid"`
	Logradouro 			Logradouro 		`gorm:"ForeignKey:endereco_codigo;AssociationForeignKey:logradouro_id"`
}

type Pessoas []Pessoa

func (p Pessoa) IsValid() map[string][]string {

	err := make(map[string][]string)

	if p.Nome == "" {
		err["nome"] = append(err["nome"], "Nome não pode ser vázio.")
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