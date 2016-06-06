package services

import (
	"academiasig-api/services/models"
)

func AuthenticationUser(user string, password string) models.Pessoa {

	var usuario models.Pessoa

	Con.Preload("Usuario").		
		Joins("JOIN usuario ON usuario.pessoa_uuid = pessoa.uuid").
		Where("usuario.login = ? AND usuario.senha = ?", user, password).
		First(&usuario)

	return usuario
}

func GetUsuarios() models.Pessoas {

	var usuarios models.Pessoas

	Con.Where("usuario_sistema = true").
		First(&usuarios)

	return usuarios
}