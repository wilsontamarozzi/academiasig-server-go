package models

type Cidade struct {
	CidadeCodigo		int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	UfCodigo			int64
	CidadeDescricao		string
	CidadeCep 			string
	Estado 				Estado 	`gorm:"foreignkey:uf_codigo;associationforeignkey:uf_codigo"`
}

type Cidades []Cidade