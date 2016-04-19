package services

import (
	"academiasig-api/services/models"
)

func AuthenticationUser(user string, password string) models.Pessoa {

	var pessoa models.Pessoa

	Con.Preload("Usuario").		
		Joins("JOIN usuario ON usuario.pessoa_id = pessoa.id").
		Where("usuario.login = ? AND usuario.senha = ?", user, password).
		First(&pessoa)

	return pessoa
}