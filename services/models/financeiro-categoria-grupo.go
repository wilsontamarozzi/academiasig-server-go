package models

import(
	"github.com/satori/go.uuid"
)

type FinanceiroCategoriaGrupo struct {
	UUID 		uuid.UUID 				`gorm:"primary_key:true"`
	Nome		string 
	Tipo		string	
	Categorias 	FinanceiroCategorias	`gorm:"ForeignKey:GrupoCategoriaUUID; AssociationForeignKey:UUID"`
}

type FinanceiroCategoriaGrupos []FinanceiroCategoriaGrupo

func (c FinanceiroCategoriaGrupo) IsValid() map[string][]string {

	err := make(map[string][]string)

	if c.Nome == "" {
		err["nome"] = append(err["nome"], "Nome não pode ser vázio.")
	}

	if c.Tipo == "" {
		err["tipo"] = append(err[""], "Tipo não pode ser vázio.")	
	}

	return err
}