package models

type CategoriaLancamento struct {
	CategoriaId 		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	CategoriaGrupoId 	int64
	Nome				string
	Tipo				string
}

type CategoriasLancamento []CategoriaLancamento