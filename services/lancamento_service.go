package services

import (
	"academiasig-api/services/models"
)

func GetLancamentos(descricao string) models.Lancamentos {

	var lancamentos models.Lancamentos

	db := Con

	if descricao != "" {
		db = db.Where("descricao iLIKE ?", "%" + descricao + "%")
	}

	db.Preload("Pessoa").
		Order("data_cadastro ASC").
		Find(&lancamentos)

    return lancamentos
}

func GetLancamentosByFullTextSearch(text string) models.Lancamentos {

	var lancamentos models.Lancamentos

	db := Con

	db.Preload("Pessoa").
		Where(`descricao iLIKE ?`, "%" + text + "%").
		Find(&lancamentos)

	return lancamentos
}

func GetLancamento(lancamentoId string) models.Lancamento {

	var lancamento models.Lancamento

	Con.Preload("Pessoa").
		Preload("CadastradoPor").
		Preload("Categorias.Categoria").
		Preload("Movimentacoes.Conta").
		Where("uuid = ?", lancamentoId).
		First(&lancamento)

	return lancamento
}

func DeleteLancamento(lancamentoId string) error {
	err := Con.Where("uuid = ?", lancamentoId).Delete(&models.Lancamento{}).Error

	return err
}

func CreateLancamento(lancamento models.Lancamento) models.Lancamento {
	
	/*
	err := Con.Create(&lancamento).Error

	if err != nil {
		panic(err)
	}
	*/

	return lancamento
}

func UpdateLancamento(lancamento models.Lancamento) error {
	err := Con.Save(&lancamento).Error

	return err
}