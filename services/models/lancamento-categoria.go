package models

import(
	"github.com/satori/go.uuid"
)

type LancamentoCategoria struct {
	UUID 					uuid.UUID 				`gorm:"primary_key:true"`
	GrupoCategoriaUUID 		uuid.UUID
	Nome					string
	Tipo					string
}

type LancamentoCategorias []LancamentoCategoria

func (c LancamentoCategoria) IsValid() map[string][]string {

	err := make(map[string][]string)

	if c.Nome == "" {
		err["nome"] = append(err["nome"], "Nome não pode ser vázio.")
	}

	if c.Tipo == "" {
		err["tipo"] = append(err[""], "Tipo não pode ser vázio.")	
	}

	return err
}