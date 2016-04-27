package services

import (
	"academiasig-api/services/models"
)

func GetContas(descricao string, tipoConta string, ativo string) models.Contas {

	descricaoQuery 	:= (map[bool]string{true: "descricao LIKE '%" + descricao + "%' AND ", false: ""}) 	[descricao != ""]
	tipoContaQuery 	:= (map[bool]string{true: "tipo_conta = '" + tipoConta + "' AND ", false: ""}) 		[tipoConta != ""]
	situacaoQuery 	:= (map[bool]string{true: "ativo = '" + ativo + "' AND ", false: ""}) 				[ativo != ""]
	
	commit := descricaoQuery + tipoContaQuery + situacaoQuery
	if commit != "" {
		commit = commit[:len(commit)-4]
	}

	var contas models.Contas

	Con.Where(commit).Find(&contas)

    return contas
}

func GetConta(contaId int64) models.Conta {

	var conta models.Conta

	Con.Preload("Banco").
		Preload("Titular").
		Preload("Operadora").
		First(&conta, contaId)

	return conta
}

func DeleteConta(contaId int64) error {
	err := Con.Where("id = ?", contaId).Delete(&models.Conta{}).Error

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