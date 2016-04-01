package services

import (
	"academiasig-api/services/models"
)

func GetGruposCategoria() models.CategoriaLancamentoGrupos {

	var gruposCategoria models.CategoriaLancamentoGrupos

	Con.Preload("CategoriasLancamento").Find(&gruposCategoria)

    return gruposCategoria
}

func GetGrupoCategoria(grupoId int64) models.CategoriaLancamentoGrupo {

	var grupoCategoria models.CategoriaLancamentoGrupo

	Con.Preload("CategoriasLancamento").First(&grupoCategoria, grupoId)

	return grupoCategoria
}

func GetGrupoCategoriaPesquisa(grupoId int64, nome string, tipo string) models.CategoriaLancamentoGrupos {

	if nome != "" {
		nome = "%" + nome + "%"
	}

	var gruposCategoria models.CategoriaLancamentoGrupos

	Con.Preload("CategoriasLancamento").
		Where("categoria_grupo_id = ?", grupoId).
		Or("nome LIKE ?", nome).
		Or("tipo = ?", tipo).
		Find(&gruposCategoria)

	return gruposCategoria
}