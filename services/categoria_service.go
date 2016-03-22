package services

import (
	"AcademiaSIG-API/services/models"
)

func GetCategorias() models.Categorias {

	var categorias models.Categorias

	Con.Find(&categorias)

    return categorias
}

func GetCategoria(categoriaId int64) models.Categoria {

	var categoria models.Categoria

	Con.First(&categoria, categoriaId)

	return categoria
}

func GetCategoriaPesquisa(categoriaId int64, categoriaGrupoId int64, nome string, tipo string) models.Categorias {

	if nome != "" {
		nome = "%" + nome + "%"
	}

	var categorias models.Categorias

	Con.Where("categoria_id = ?", categoriaId).
			Or("categoria_grupo_id = ?", categoriaGrupoId).
			Or("nome LIKE ?", nome).
			Or("tipo = ?", tipo).
		Find(&categorias)

	return categorias
}

func SaveCategoria(categoriaGrupoId int64, nome string) models.Categoria {

	var grupoCategoria models.CategoriaLancamentoGrupo

	Con.First(&grupoCategoria, categoriaGrupoId)

	categoria := models.CategoriaLancamento{
		CategoriaGrupoId: categoriaGrupoId,
		Nome: nome,
		Tipo: grupoCategoria.Tipo,
	}

	Con.Create(&categoria)

	return categoria
}