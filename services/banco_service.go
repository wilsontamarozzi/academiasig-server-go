package services

import (
	"academiasig-api/services/models"
)

func GetBancos(nome string, numero string) models.Bancos {

	var bancos models.Bancos

	db := Con

	db = db.Table("banco").Select("*")

	if nome != "" {
		db = db.Where("nome iLIKE ?", "%" + nome + "%")
	}

	if numero != "" {
		db = db.Where("numero iLIKE ?", "%" + numero + "%")
	}

	db.Find(&bancos)

    return bancos
}

func GetBancosByFullTextSearch(text string) models.Bancos {

	var bancos models.Bancos

	db := Con

	db.Where(`
		nome || ' ' || 
		numero iLIKE ?`, "%" + text + "%").
	Find(&bancos)

	return bancos
}

func GetBanco(bancoId string) models.Banco {

	var banco models.Banco

	Con.Where("uuid = ?", bancoId).
		First(&banco)

	return banco
}

func DeleteBanco(bancoId string) error {
	err := Con.Where("uuid = ?", bancoId).Delete(&models.Banco{}).Error

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