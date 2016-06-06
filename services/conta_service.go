package services

import (
	"academiasig-api/services/models"
)

func GetContas(descricao string, tipoConta string, ativo string) models.Contas {

	var contas models.Contas

	db := Con

	if descricao != "" {
		db = db.Where("descricao iLIKE ?", "%" + descricao + "%")
	}

	if tipoConta != "" {
		db = db.Where("tipo_conta iLIKE ?", "%" + tipoConta + "%")	
	}

	if ativo != "" {
		db = db.Where("ativo = ?", ativo)
	}
 
	db.Find(&contas)

    return contas
}

func GetContasByFullTextSearch(text string, tipoConta string, ativo string) models.Contas {

	var contas models.Contas

	db := Con

	if tipoConta != "" {
		db = db.Where("tipo_conta = ?", tipoConta)
	}

	if ativo != "" {
		db = db.Where("ativo = ?", ativo)
	}

	db.Where(`descricao iLIKE ?`, "%" + text + "%").Find(&contas)

	return contas
}

func GetConta(contaId string) models.Conta {

	var conta models.Conta

	Con.Preload("Banco").
		Preload("Titular").
		Preload("Operadora").
		Where("uuid = ?", contaId).
		First(&conta)

	return conta
}

func DeleteConta(contaId string) error {
	err := Con.Where("uuid = ?", contaId).Delete(&models.Conta{}).Error

	return err
}

func CreateConta(conta models.Conta) models.Conta {
	err := Con.Set("gorm:save_associations", false).Create(&conta).Error

	if err != nil {
		panic(err)
	}

	return conta
}

func UpdateConta(conta models.Conta) error {
	err := Con.Set("gorm:save_associations", false).Save(&conta).Error

	return err
}