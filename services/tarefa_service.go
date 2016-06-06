package services

import (
	"github.com/jinzhu/gorm"
	"academiasig-api/services/models"
)

func GetTarefas(descricao string, concluida string, responsavelUUID string, categoriaUUID string, tipoData string, dataInicio string, dataFim string) models.Tarefas {

	var tarefas models.Tarefas
	var tipoDataQuery string


	db := Con

	if descricao != "" {
		db = db.Where("descricao iLIKE ?", "%" + descricao + "%")
	}

	if concluida != "" {
		db = db.Where("concluida = ?", concluida)
	}

	if responsavelUUID != "" {
		db = db.Where("responsavel_uuid = ?", responsavelUUID)
	}

	if categoriaUUID != "" {
		db = db.Where("categoria_uuid = ?", categoriaUUID)
	}

	if tipoData != "" {
		switch tipoData {
			case "cadastro": tipoDataQuery = "data_cadastro"
			case "vencimento": tipoDataQuery = "data_vencimento"
			case "conclusao": tipoDataQuery = "data_conclusao"
			default: tipoDataQuery = "data_cadastro"
		}
	}

	if dataInicio != "" && dataFim != "" {
		db = db.Where(tipoDataQuery + " BETWEEN ? AND ?", dataInicio, dataFim)
	} else if dataInicio != "" {
		db = db.Where(tipoDataQuery + " > ?", dataInicio)
	} else if dataFim != "" {
		db = db.Where(tipoDataQuery + " < ?", dataFim)
	}
	
	db.Preload("Categoria").
		Preload("Responsavel").
		Order("data_cadastro ASC").
		Find(&tarefas)

    return tarefas
}

func GetTarefasByFullTextSearch(text string, concluida string, responsavelUUID string, categoriaUUID string, tipoData string, dataInicio string, dataFim string) models.Tarefas {

	var tarefas models.Tarefas
	var tipoDataQuery string

	db := Con

	if concluida != "" {
		db = db.Where("concluida = ?", concluida)
	}

	if responsavelUUID != "" {
		db = db.Where("responsavel_uuid = ?", responsavelUUID)
	}

	if categoriaUUID != "" {
		db = db.Where("categoria_uuid = ?", categoriaUUID)
	}

	if tipoData != "" {
		switch tipoData {
			case "cadastro": tipoDataQuery = "data_cadastro"
			case "vencimento": tipoDataQuery = "data_vencimento"
			case "conclusao": tipoDataQuery = "data_conclusao"
			default: tipoDataQuery = "data_cadastro"
		}
	}

	if dataInicio != "" && dataFim != "" {
		db = db.Where(tipoDataQuery + " BETWEEN ? AND ?", dataInicio, dataFim)
	} else if dataInicio != "" {
		db = db.Where(tipoDataQuery + " > ?", dataInicio)
	} else if dataFim != "" {
		db = db.Where(tipoDataQuery + " < ?", dataFim)
	}

	db.Preload("Categoria").
		Preload("Responsavel").
		Where(`descricao iLIKE ?`, "%" + text + "%").
		Find(&tarefas)

	return tarefas
}

func GetTarefa(tarefaId string) models.Tarefa {

	var tarefa models.Tarefa

	Con.Preload("Categoria").
		Preload("Responsavel").
		Preload("Comentarios", func(db *gorm.DB) *gorm.DB {
		    return db.Order("tarefa_comentario.data_cadastro DESC")
		}).
		Where("uuid = ?", tarefaId).
		First(&tarefa)

	return tarefa
}

func DeleteTarefa(tarefaId string) error {
	err := Con.Where("uuid = ?", tarefaId).Delete(&models.Tarefa{}).Error

	return err
}

func CreateTarefa(tarefa models.Tarefa) models.Tarefa {
	err := Con.Set("gorm:save_associations", true).Create(&tarefa).Error

	if err != nil {
		panic(err)
	}

	return tarefa
}

func UpdateTarefa(tarefa models.Tarefa) error {
	err := Con.Set("gorm:save_associations", true).Save(&tarefa).Error

	return err
}