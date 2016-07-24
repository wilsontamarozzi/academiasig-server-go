package models

import(
	"time"
	"github.com/satori/go.uuid"
)

type LancamentoMovimentacao struct {
	UUID 				uuid.UUID 	`gorm:"primary_key:true"`
	LancamentoUUID 		uuid.UUID
	ContaUUID 			uuid.UUID
	Valor				float64
	DataMovimentacao 	time.Time

	Conta 				Conta 		`gorm:"ForeignKey:ContaUUID; AssociationForeignKey:UUID"`
}

type LancamentoMovimentacoes []LancamentoMovimentacao

func (c LancamentoMovimentacao) IsValid() map[string][]string {

	err := make(map[string][]string)

	return err
}