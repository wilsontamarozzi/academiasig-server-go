package models

type Conta struct {
	ContaId 			int64		`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	BancoId				int64 
	TitularId			int64 
	OperadoraId			int64
	TipoConta			string
	Descricao			string
	SaldoAgencia		int64
	Agencia 			int64
	Digito				int64
	VencimentoFatura	int64
	FechamentoFatura 	int64
	Ativo 				bool
	Banco 				*Banco 		`gorm:"foreignkey:banco_id;associationforeignkey:banco_id"`
	Titular 			*Pessoa 	`gorm:"foreignkey:titular_id;associationforeignkey:pessoa_id"`
	Operadora 			*Pessoa 	`gorm:"foreignkey:operadora_id;associationforeignkey:pessoa_id"`
}

type Contas []Conta