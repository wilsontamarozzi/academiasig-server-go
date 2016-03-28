package services

import (
	"AcademiaSIG-API/services/models"
)

func AuthenticationUser(user string, password string) bool {

	var usuario models.Usuario
	var count int

	Con.Model(&usuario).Where("login = ? AND senha = ?", user, password).Count(&count)

	if count == 1 {
		return true
	} else {
		return false
	}
}