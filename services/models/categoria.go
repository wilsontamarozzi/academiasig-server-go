package models

type CategoriaLancamento struct {
	Id 					int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	CategoriaGrupoId 	int64
	Nome				string
	Tipo				string
}

type CategoriasLancamento []CategoriaLancamento