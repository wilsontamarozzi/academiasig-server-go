package services

import (
	"academiasig-api/services/models"
)

func GetLancamentoCategorias(lancamentoUUID string) models.LancamentoCategorias {

	var lancamentoCategorias models.LancamentoCategorias

	db := Con

	if lancamentoUUID != "" {
		db = db.Where("lancamento_uudi = ?", lancamentoUUID)
	}

	db.Find(&lancamentoCategorias)

    return lancamentoCategorias
}

func GetLancamentoCategoria(lancamentoCategoriaId string) models.LancamentoCategoria {

	var lancamentoCategoria models.LancamentoCategoria

	Con.Where("uuid = ?", lancamentoCategoriaId).
		First(&lancamentoCategoria)

	return lancamentoCategoria
}

func DeleteLancamentoCategoria(lancamentoCategoriaId string) error {
	err := Con.Where("uuid = ?", lancamentoCategoriaId).Delete(&models.LancamentoCategoria{}).Error

	return err
}

func CreateLancamentoCategoria(lancamentoCategoria models.LancamentoCategoria) models.LancamentoCategoria {
	err := Con.Set("gorm:save_associations", false).Create(&lancamentoCategoria).Error

	if err != nil {
		panic(err)
	}

	return lancamentoCategoria
}

func UpdateLancamentoCategoria(lancamentoCategoria models.LancamentoCategoria) error {
	err := Con.Set("gorm:save_associations", false).Save(&lancamentoCategoria).Error

	return err
}