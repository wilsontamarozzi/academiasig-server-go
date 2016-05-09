package database

import (
	"github.com/jinzhu/gorm"
  "academiasig-api/services/models"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetConnection() *gorm.DB {
  //MYSQL
  //db, err := gorm.Open("mysql", "academiasigapi:4321@tcp(wilsontamarozzi.dyndns.org:3306)/academiasig-api?charset=utf8&parseTime=True")
	
  //POSTGRESQL - Localhost
  //db, err := gorm.Open("postgres", "host=localhost user=academiasigapi dbname=academiasig-api sslmode=disable password=4321")

  //POSTGRESQL - Heroku
  db, err := gorm.Open("postgres", "host=ec2-54-235-199-36.compute-1.amazonaws.com user=jwgsljcoyxoxag dbname=dfoq7q8so3lv6n password=SdFlEyF_SsGqxRtPTIC6HlQZ_X")

	if err != nil {
      panic(err)
    	//panic("Houve uma falha ao conectar ao banco de dados")
  }

  //Faz com que o nome dos objetos fique no singular
  db.SingularTable(true)
  //Ativa log de todas as saidas da conexão (SQL)
  db.LogMode(false)
  //Seta o maximo de conexões
  db.DB().SetMaxIdleConns(1)
  db.DB().SetMaxOpenConns(1)

  //autoMigrate(db)
  //autoPopulate(db)

 	return db
}

func autoMigrate(db *gorm.DB) {
  //db.AutoMigrate(&models.Pessoa{}, &models.Usuario{}, &models.Pais{}, &models.Estado{}, &models.Bairro{}, &models.Cidade{}, &models.Logradouro{}, &models.Banco{}, &models.Conta{})
  //db.AutoMigrate(&models.Pais{})
}

func autoPopulate(db *gorm.DB) {

  var pessoa models.Pessoa

  pessoa.Nome = "Wilson"
  pessoa.TipoPessoa = "F"
  pessoa.Ativo = true
  pessoa.UsuarioSistema = true

  db.Set("gorm:save_associations", false).Create(&pessoa)

  var usuario models.Usuario

  usuario.PessoaId = 1
  usuario.Ativo = true
  usuario.Login = "wilson"
  usuario.Senha = "202cb962ac59075b964b07152d234b70"

  db.Set("gorm:save_associations", false).Create(&usuario)
}