package services

import (
	"academiasig-api/services/models"
)

func GetTarefaCategorias(descricao string) models.TarefaCategorias {

	var tarefaCategorias models.TarefaCategorias

	db := Con

	if descricao != "" {
		db = db.Where("descricao iLIKE ?", "%" + descricao + "%")
	}

	db.Order("descricao ASC").Find(&tarefaCategorias)

    return tarefaCategorias
}

func GetTarefaCategoriasByFullTextSearch(text string) models.TarefaCategorias {

	var tarefaCategorias models.TarefaCategorias

	db := Con

	db.Where(`descricao iLIKE ?`, "%" + text + "%").Find(&tarefaCategorias)

	return tarefaCategorias
}

func GetTarefaCategoria(tarefaCategoriaId string) models.TarefaCategoria {

	var tarefaCategoria models.TarefaCategoria

	Con.Where("uuid = ?", tarefaCategoriaId).
		First(&tarefaCategoria)

	return tarefaCategoria
}

func DeleteTarefaCategoria(tarefaCategoriaId string) error {
	err := Con.Where("uuid = ?", tarefaCategoriaId).Delete(&models.TarefaCategoria{}).Error

	return err
}

func CreateTarefaCategoria(tarefaCategoria models.TarefaCategoria) models.TarefaCategoria {
	err := Con.Create(&tarefaCategoria).Error

	if err != nil {
		panic(err)
	}

	return tarefaCategoria
}

func UpdateTarefaCategoria(tarefaCategoria models.TarefaCategoria) error {
	err := Con.Save(&tarefaCategoria).Error

	return err
}