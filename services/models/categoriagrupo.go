package models

type CategoriaLancamentoGrupo struct {
	CategoriaGrupoId 	int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT" json:"categoriaGrupoId"`
	Nome				string 	`json:"nome"`
	Tipo				string 	`json:"tipo"`
	CategoriasLancamento CategoriasLancamento `gorm:"foreignkey:categoria_grupo_id;associationforeignkey:categoria_grupo_id"`
}

type CategoriaLancamentoGrupos []CategoriaLancamentoGrupo