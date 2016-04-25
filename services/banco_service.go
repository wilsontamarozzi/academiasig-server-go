package services

import (
	"academiasig-api/services/models"
)

func GetBancos(nome string, numero string) models.Bancos {

	nomeQuery := (map[bool]string{true: "nome LIKE '%" + nome + "%' AND ", false: ""}) [nome != ""]
	numeroQuery := (map[bool]string{true: "numero LIKE '%" + numero + "%' AND ", false: ""}) [numero != ""]
	
	commit := nomeQuery + numeroQuery
	if commit != "" {
		commit = commit[:len(commit)-4]
	}

	var bancos models.Bancos

	Con.Where(commit).Find(&bancos)

    return bancos
}

func GetBanco(bancoId int64) models.Banco {

	var banco models.Banco

	Con.First(&banco, bancoId)

	return banco
}

func DeleteBanco(bancoId int64) error {
	err := Con.Where("id = ?", bancoId).Delete(&models.Banco{}).Error

	return err
}

func CreateBanco(banco models.Banco) models.Banco {
	err := Con.Create(&banco).Error

	if err != nil {
		panic(err)
	}

	return banco
}

func UpdateBanco(banco models.Banco) error {
	err := Con.Save(&banco).Error

	return err
}