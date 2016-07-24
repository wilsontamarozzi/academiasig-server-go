package services

import (
	"github.com/jinzhu/gorm"
	"academiasig-api/services/models"
)

func GetGruposCategoria(nome string) models.FinanceiroCategoriaGrupos {

	var gruposCategoria models.FinanceiroCategoriaGrupos

	db := Con

	if nome != "" {
		db = db.Where("nome iLIKE ?", "%" + nome + "%")
	}

	db.Preload("Categorias", func(db *gorm.DB) *gorm.DB {
		    return db.Order("financeiro_categoria.nome ASC")
		}).
		Order("tipo ASC, nome ASC").
		Find(&gruposCategoria)

    return gruposCategoria
}

func GetGruposCategoriaByFullTextSearch(text string) models.FinanceiroCategoriaGrupos {

	var gruposCategoria models.FinanceiroCategoriaGrupos

	db := Con

	db.Where(`nome iLIKE ?`, "%" + text + "%").
		Preload("Categorias").
		Find(&gruposCategoria)

	return gruposCategoria
}

func GetGrupoCategoria(grupoCategoriaId string) models.FinanceiroCategoriaGrupo {

	var grupoCategoria models.FinanceiroCategoriaGrupo

	Con.Where("uuid = ?", grupoCategoriaId).
		Preload("Categorias").
		First(&grupoCategoria)

	return grupoCategoria
}

func DeleteGrupoCategoria(grupoCategoriaId string) error {
	err := Con.Where("uuid = ?", grupoCategoriaId).Delete(&models.FinanceiroCategoriaGrupo{}).Error

	return err
}

func CreateGrupoCategoria(grupoCategoria models.FinanceiroCategoriaGrupo) models.FinanceiroCategoriaGrupo {
	err := Con.Create(&grupoCategoria).Error

	if err != nil {
		panic(err)
	}

	return grupoCategoria
}

func UpdateGrupoCategoria(grupoCategoria models.FinanceiroCategoriaGrupo) error {
	err := Con.Save(&grupoCategoria).Error

	return err
}