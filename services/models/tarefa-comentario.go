package models

import(
	"time"
	"github.com/satori/go.uuid"
)

type TarefaComentario struct {
	UUID 			uuid.UUID 	`gorm:"primary_key:true"`
	TarefaUUID 		uuid.UUID
	DataCadastro	time.Time 	`gorm:"default:'NOW()'"`
	Comentario 		string
}

type TarefaComentarios []TarefaComentario

func (t TarefaComentario) IsValid() map[string][]string {

	err := make(map[string][]string)

	if t.Comentario == "" {
		err["comentario"] = append(err["comentario"], "Comentário não pode ser vázio.")
	}

	return err
}