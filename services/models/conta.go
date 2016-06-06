package models

import(
	"github.com/satori/go.uuid"
)

type Conta struct {
	UUID 				uuid.UUID		`gorm:"primary_key:true"`
	TipoConta			string
	Ativo 				*bool
	Descricao			string
	Agencia 			int64
	AgenciaDigito		int64
	Conta 				int64
	ContaDigito 		int64
	
	VencimentoFatura	int64
	FechamentoFatura 	int64

	SaldoInicial		float64
	
	BancoUUID			uuid.UUID 
	TitularUUID			uuid.UUID
	OperadoraUUID		uuid.UUID

	Banco 				*Banco 		`gorm:"foreignkey:banco_uuid;associationforeignkey:uuid"`
	Titular 			*Pessoa 	`gorm:"foreignkey:titular_uuid;associationforeignkey:uuid"`
	Operadora 			*Pessoa 	`gorm:"foreignkey:operadora_uuid;associationforeignkey:uuid"`
}

type Contas []Conta

func (c Conta) IsValid() map[string][]string {

	err := make(map[string][]string)

	if c.Descricao == "" {
		err["descricao"] = append(err["descricao"], "Descricao não pode ser vázio.")
	}

	if len(c.Descricao) < 2 {
		err["descricao"] = append(err["descricao"], "Descricao deve ter mais de 2 caracteres.")
	}

	if c.TipoConta == "" {
		err["tipoConta"] = append(err["tipoConta"], "Tipo de conta não definido.")	
	} else {
		if c.TipoConta == "Conta Corrente" {
			if c.Agencia <= 0 {
				err["agencia"] = append(err["agencia"], "Agência não pode estar vázio.")	
			}

			if c.Conta <= 0 {
				err["conta"] = append(err["conta"], "Conta não pode estar vázio.")
			}

			/*
			if len(c.BancoUUID) < 0 {
				err["banco"] = append(err["banco"], "Informe um banco.")
			}
			*/	
		}
	}

	if c.Ativo == nil {
		err["ativo"] = append(err["ativo"], "Campo Ativo não pode ser nulo.")
	}

	return err
}