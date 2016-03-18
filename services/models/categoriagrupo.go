package models

type CategoriaLancamentoGrupo struct {
	CategoriaGrupoId 	int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT"`
	Nome				string 
	Tipo				string	
	CategoriasLancamento CategoriasLancamento `gorm:"foreignkey:categoria_grupo_id;associationforeignkey:categoria_grupo_id"`
}

type CategoriaLancamentoGrupos []CategoriaLancamentoGrupo