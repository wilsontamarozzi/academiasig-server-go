package services

import (
	"academiasig-api/services/models"
)

func GetPessoas(pessoaId string, ativo string, nome string, email string, tipoPessoa string, usuarioSistema string) models.Pessoas {

	var pessoas models.Pessoas

	db := Con

	if pessoaId != "" {
		db = db.Where("uuid = ?", pessoaId)
	}

	if ativo != "" {
		db = db.Where("ativo = ?", ativo)
	}

	if nome != "" {
		db = db.Where("nome iLIKE ?", "%" + nome + "%")	
	}

	if email != "" {
		db = db.Where("email iLIKE ?", "%" + email + "%")
	}

	if usuarioSistema != "" {
		db = db.Where("usuario_sistema = ?", usuarioSistema)
	}

	if tipoPessoa != "" {
		db = db.Where("tipo_pessoa = ?", tipoPessoa)
	}
	
	db.Find(&pessoas)

    return pessoas
}

func GetPessoasByFullTextSearch(text string, tipoPessoa string, ativo string) models.Pessoas {

	var pessoas models.Pessoas

	db := Con

	if tipoPessoa != "" {
		db = db.Where("tipo_pessoa = ?", tipoPessoa)
	}

	if ativo != "" {
		db = db.Where("ativo = ?", ativo)
	}

	db.Where(`
		nome || ' ' || 
		email || ' ' || 
		observacao || ' ' || 
		observacao || ' ' || 
		suporte || ' ' || 
		complemento || ' ' || 
		cpf || ' ' || 
		rg || ' ' || 
		telefone_celular || ' ' || 
		telefone_empresa || ' ' || 
		telefone_residencial || ' ' || 
		cnpj || ' ' || 
		inscricao_estadual || ' ' || 
		inscricao_municipal || ' ' || 
		razao_social || ' ' || 
		telefone_comercial || ' ' || 
		website iLIKE ?`, "%" + text + "%").
	Find(&pessoas)

	return pessoas
}

func GetPessoa(pessoaId string) models.Pessoa {

	var pessoa models.Pessoa

	Con.Preload("Logradouro.Bairro.Cidade.Estado").
		Preload("Usuario").
		Where("uuid = ?", pessoaId).
		First(&pessoa)

	return pessoa
}

func DeletePessoa(pessoaId string) error {
	err := Con.Where("uuid = ?", pessoaId).Delete(&models.Pessoa{}).Error

	return err
}

func CreatePessoa(pessoa models.Pessoa) models.Pessoa {
	err := Con.Set("gorm:save_associations", false).Create(&pessoa).Error

	if err != nil {
		panic(err)
	}

	return pessoa
}

func UpdatePessoa(pessoa models.Pessoa) error {
	err := Con.Set("gorm:save_associations", false).Save(&pessoa).Error

	return err
}