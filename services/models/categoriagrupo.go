package models

type CategoriaLancamentoGrupo struct {
	Id 						int64					`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	Nome					string 
	Tipo					string	
	CategoriasLancamento 	CategoriasLancamento	`gorm:"ForeignKey:CategoriaGrupoId; AssociationForeignKey:CategoriaGrupoId"`
}

type CategoriaLancamentoGrupos []CategoriaLancamentoGrupo