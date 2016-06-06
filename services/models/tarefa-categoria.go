package models

import(
	"github.com/satori/go.uuid"
)

type TarefaCategoria struct {
	UUID 		uuid.UUID `gorm:"primary_key:true"`
	Descricao 	string
}

type TarefaCategorias []TarefaCategoria

func (t TarefaCategoria) IsValid() map[string][]string {

	err := make(map[string][]string)

	if t.Descricao == "" {
		err["descricao"] = append(err["descricao"], "Descrição não pode ser vázio.")
	}

	return err
}