package services

import (
	"academiasig-api/services/models"
)

func GetCategorias(nome string) models.FinanceiroCategorias {

	var categorias models.FinanceiroCategorias

	db := Con

	if nome != "" {
		db = db.Where("nome iLIKE ?", "%" + nome + "%")
	}

	db.Find(&categorias)

    return categorias
}

func GetCategoriasByFullTextSearch(text string) models.FinanceiroCategorias {

	var categorias models.FinanceiroCategorias

	db := Con

	db.Where(`nome iLIKE ?`, "%" + text + "%").
		Find(&categorias)

	return categorias
}

func GetCategoria(categoriaId string) models.FinanceiroCategoria {

	var categoria models.FinanceiroCategoria

	Con.Where("uuid = ?", categoriaId).
		First(&categoria)

	return categoria
}

func DeleteCategoria(categoriaId string) error {
	err := Con.Where("uuid = ?", categoriaId).Delete(&models.FinanceiroCategoria{}).Error

	return err
}

func CreateCategoria(categoria models.FinanceiroCategoria) models.FinanceiroCategoria {
	err := Con.Create(&categoria).Error

	if err != nil {
		panic(err)
	}

	return categoria
}

func UpdateCategoria(categoria models.FinanceiroCategoria) error {
	err := Con.Save(&categoria).Error

	return err
}