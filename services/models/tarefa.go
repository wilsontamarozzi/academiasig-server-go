package models

import(
	"time"
	"github.com/satori/go.uuid"
)

type Tarefa struct {
	UUID 				uuid.UUID 			`gorm:"primary_key:true"`
	Descricao 			string
	DataCadastro 		time.Time 			`gorm:"default:'NOW()'"`
	DataVencimento 		time.Time 			`gorm:"default:'NOW()'"`
	DataConclusao 		*time.Time
	Concluida 			bool

	CategoriaUUID 		uuid.UUID 
	ResponsavelUUID 	uuid.UUID

	Categoria 			TarefaCategoria		`gorm:"ForeignKey:categoria_uuid;"`
	Responsavel 		Pessoa 				`gorm:"ForeignKey:responsavel_uuid;"`
	Comentarios			TarefaComentarios 	`gorm:"ForeignKey:tarefa_uuid;AssociationForeignKey:uuid"`
}

type Tarefas []Tarefa

func (t Tarefa) IsValid() map[string][]string {

	err := make(map[string][]string)

	if t.Descricao == "" {
		err["descricao"] = append(err["descricao"], "Descrição não pode ser vázio.")
	}

	return err
}