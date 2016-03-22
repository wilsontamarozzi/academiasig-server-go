package services

import (
	"AcademiaSIG-API/services/models"
)

func GetContas() models.Contas {

	var contas models.Contas

	Con.Preload("Banco").
		Preload("Titular").
		Preload("Operadora").
		Find(&contas)

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

func GetContaPesquisa(contaId int64, descricao string) models.Contas {

	if descricao != "" {
		descricao = "%" + descricao + "%"
	}

	var contas models.Contas

	Con.Preload("Banco").
		Preload("Titular").
		Preload("Operadora").
		Where("conta_id = ?", contaId).
			Or("descricao LIKE ?", descricao).
		Find(&contas)

	return contas
}