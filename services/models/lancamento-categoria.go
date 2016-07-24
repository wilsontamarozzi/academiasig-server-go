package models

import(
	"github.com/satori/go.uuid"
)

type LancamentoCategoria struct {
	UUID 					uuid.UUID 			`gorm:"primary_key:true"`
	LancamentoUUID 			uuid.UUID
	FinanceiroCategoriaUUID uuid.UUID
	Valor					float64

	Categoria 				FinanceiroCategoria `gorm:"ForeignKey:FinanceiroCategoriaUUID; AssociationForeignKey:UUID"`
}

type LancamentoCategorias []LancamentoCategoria

func (c LancamentoCategoria) IsValid() map[string][]string {

	err := make(map[string][]string)

	return err
}