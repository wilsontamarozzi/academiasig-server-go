package models

type Conta struct {
	Id 					int64		`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
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
	
	BancoId				*int64 
	TitularId			*int64 
	OperadoraId			*int64

	Banco 				*Banco 		`gorm:"foreignkey:banco_id;associationforeignkey:id"`
	Titular 			*Pessoa 	`gorm:"foreignkey:titular_id;associationforeignkey:id"`
	Operadora 			*Pessoa 	`gorm:"foreignkey:operadora_id;associationforeignkey:id"`
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

			if c.BancoId == nil {
				err["banco"] = append(err["banco"], "Informe um banco.")
			}
		}
	}

	if c.Ativo == nil {
		err["ativo"] = append(err["ativo"], "Campo Ativo não pode ser nulo.")
	}

	return err
}