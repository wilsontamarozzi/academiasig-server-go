package categoriagrupo

type GrupoCategoria struct {
	CategoriaGrupoId 	int64	`gorm:"column:categoria_grupo_id; primary_key;" sql:"AUTO_INCREMENT" json:"categoriaGrupoId"`
	Nome				string 	`json:"nome"`
	Tipo				string 	`json:"tipo"`
}

type GruposCategoria []GrupoCategoria