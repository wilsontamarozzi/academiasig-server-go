package services

import (
	"github.com/jinzhu/gorm"
	"academiasig-api/services/models"
)

func GetGruposCategoria(nome string) models.LancamentoCategoriaGrupos {

	var gruposCategoria models.LancamentoCategoriaGrupos

	db := Con

	if nome != "" {
		db = db.Where("nome iLIKE ?", "%" + nome + "%")
	}

	db.Preload("Categorias", func(db *gorm.DB) *gorm.DB {
		    return db.Order("lancamento_categoria.nome ASC")
		}).
		Order("tipo ASC, nome ASC").
		Find(&gruposCategoria)

    return gruposCategoria
}

func GetGruposCategoriaByFullTextSearch(text string) models.LancamentoCategoriaGrupos {

	var gruposCategoria models.LancamentoCategoriaGrupos

	db := Con

	db.Where(`nome iLIKE ?`, "%" + text + "%").
		Preload("Categorias").
		Find(&gruposCategoria)

	return gruposCategoria
}

func GetGrupoCategoria(grupoCategoriaId string) models.LancamentoCategoriaGrupo {

	var grupoCategoria models.LancamentoCategoriaGrupo

	Con.Where("uuid = ?", grupoCategoriaId).
		Preload("Categorias").
		First(&grupoCategoria)

	return grupoCategoria
}

func DeleteGrupoCategoria(grupoCategoriaId string) error {
	err := Con.Where("uuid = ?", grupoCategoriaId).Delete(&models.LancamentoCategoriaGrupo{}).Error

	return err
}

func CreateGrupoCategoria(grupoCategoria models.LancamentoCategoriaGrupo) models.LancamentoCategoriaGrupo {
	err := Con.Create(&grupoCategoria).Error

	if err != nil {
		panic(err)
	}

	return grupoCategoria
}

func UpdateGrupoCategoria(grupoCategoria models.LancamentoCategoriaGrupo) error {
	err := Con.Save(&grupoCategoria).Error

	return err
}