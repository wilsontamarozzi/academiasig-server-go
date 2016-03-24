package services

import (
	"AcademiaSIG-API/services/models"
)

func GetCategorias() models.CategoriaLancamento {

	var categorias models.CategoriaLancamento

	Con.Find(&categorias)

    return categorias
}

func GetCategoria(categoriaId int64) models.CategoriaLancamento {

	var categoria models.CategoriaLancamento

	Con.First(&categoria, categoriaId)

	return categoria
}

func GetCategoriaPesquisa(categoriaId int64, categoriaGrupoId int64, nome string, tipo string) models.CategoriaLancamento {

	if nome != "" {
		nome = "%" + nome + "%"
	}

	var categorias models.CategoriaLancamento

	Con.Where("categoria_id = ?", categoriaId).
			Or("categoria_grupo_id = ?", categoriaGrupoId).
			Or("nome LIKE ?", nome).
			Or("tipo = ?", tipo).
		Find(&categorias)

	return categorias
}

func SaveCategoria(categoriaGrupoId int64, nome string) models.CategoriaLancamento {

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