package services

import (
	"academiasig-api/services/models"
)

func GetLancamentoMovimentacaos(lancamentoUUID string) models.LancamentoMovimentacoes {

	var lancamentoMovimentacoes models.LancamentoMovimentacoes

	db := Con

	if lancamentoUUID != "" {
		db = db.Where("lancamento_uudi = ?", lancamentoUUID)
	}

	db.Find(&lancamentoMovimentacoes)

    return lancamentoMovimentacoes
}

func GetLancamentoMovimentacao(lancamentoMovimentacaoId string) models.LancamentoMovimentacao {

	var lancamentoMovimentacao models.LancamentoMovimentacao

	Con.Where("uuid = ?", lancamentoMovimentacaoId).
		First(&lancamentoMovimentacao)

	return lancamentoMovimentacao
}

func DeleteLancamentoMovimentacao(lancamentoMovimentacaoId string) error {
	err := Con.Where("uuid = ?", lancamentoMovimentacaoId).Delete(&models.LancamentoMovimentacao{}).Error

	return err
}

func CreateLancamentoMovimentacao(lancamentoMovimentacao models.LancamentoMovimentacao) models.LancamentoMovimentacao {
	err := Con.Set("gorm:save_associations", false).Create(&lancamentoMovimentacao).Error

	if err != nil {
		panic(err)
	}

	return lancamentoMovimentacao
}

func UpdateLancamentoMovimentacao(lancamentoMovimentacao models.LancamentoMovimentacao) error {
	err := Con.Set("gorm:save_associations", false).Save(&lancamentoMovimentacao).Error

	return err
}