package models

type Conta struct {
	ContaId 			int64		`gorm:"column:banco_id; primary_key;" sql:"AUTO_INCREMENT"`
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
	Titular 			*Pessoa 	`gorm:"foreignkey:pessoa_id;associationforeignkey:pessoa_id"`
	Operadora 			*Pessoa 	`gorm:"foreignkey:operadora_id;associationforeignkey:operadora_id"`
}

type Contas []Conta