package services

import (
	"academiasig-api/services/models"
)

func GetCategorias(nome string) models.LancamentoCategorias {

	var categorias models.LancamentoCategorias

	db := Con

	if nome != "" {
		db = db.Where("nome iLIKE ?", "%" + nome + "%")
	}

	db.Find(&categorias)

    return categorias
}

func GetCategoriasByFullTextSearch(text string) models.LancamentoCategorias {

	var categorias models.LancamentoCategorias

	db := Con

	db.Where(`nome iLIKE ?`, "%" + text + "%").
		Find(&categorias)

	return categorias
}

func GetCategoria(categoriaId string) models.LancamentoCategoria {

	var categoria models.LancamentoCategoria

	Con.Where("uuid = ?", categoriaId).
		First(&categoria)

	return categoria
}

func DeleteCategoria(categoriaId string) error {
	err := Con.Where("uuid = ?", categoriaId).Delete(&models.LancamentoCategoria{}).Error

	return err
}

func CreateCategoria(categoria models.LancamentoCategoria) models.LancamentoCategoria {
	err := Con.Create(&categoria).Error

	if err != nil {
		panic(err)
	}

	return categoria
}

func UpdateCategoria(categoria models.LancamentoCategoria) error {
	err := Con.Save(&categoria).Error

	return err
}