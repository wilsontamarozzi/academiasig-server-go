package categoriagrupo

type CategoriaLancamentoGrupo struct {
	CategoriaGrupoId 	int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT" json:"categoriaGrupoId"`
	Nome				string 	`json:"nome"`
	Tipo				string 	`json:"tipo"`
	CategoriasLancamento CategoriasLancamento `gorm:"foreignkey:categoria_grupo_id;associationforeignkey:categoria_grupo_id"`
}

type CategoriaLancamentoGrupos []CategoriaLancamentoGrupo

type CategoriaLancamento struct {
	CategoriaId 		int64	`gorm:"primary_key;" sql:"AUTO_INCREMENT" json:"categoriaId"`
	CategoriaGrupoId 	int64	`gorm:"index" json:"categoriaGrupoId"`
	Nome				string 	`json:"nome"`
	Tipo				string 	`json:"tipo"`
}

type CategoriasLancamento []CategoriaLancamento