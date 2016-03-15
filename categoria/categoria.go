package categoria

type Categoria struct {
	CategoriaId 		int64	`gorm:"column:categoria_id; primary_key;" sql:"AUTO_INCREMENT" json:"categoriaId"`
	CategoriaGrupoId 	int64	`gorm:"column:categoria_grupo_id" json:"categoriaGrupoId"`
	Nome				string 	`json:"nome"`
	Tipo				string 	`json:"tipo"`
}

type Categorias []Categoria