package database

import (
	//"fmt"
	//"reflect"

	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:02371890@tcp(192.168.0.13:3306)/sisacademia2?charset=utf8&parseTime=True")
	
	if err != nil {
    	panic("Houve uma falha ao conectar ao banco de dados")
  		//fmt.Println("Houve uma falha ao conectar ao banco de dados")
  	}

  	//fmt.Println(reflect.TypeOf(db))
  	return db
}