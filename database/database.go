package database

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
  db, err := gorm.Open("mysql", "academiasigapi:4321@tcp(wilsontamarozzi.dyndns.org:3306)/academiasig-api?charset=utf8&parseTime=True")
	
	if err != nil {
    	panic("Houve uma falha ao conectar ao banco de dados")
  }

  //Faz com que o nome dos objetos fique no singular
  db.SingularTable(true)
  //Ativa log de todas as saidas da conexão (SQL)
  db.LogMode(false)
  //Seta o maximo de conexões
  db.DB().SetMaxIdleConns(1)
  db.DB().SetMaxOpenConns(1)

 	return db
}