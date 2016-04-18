package models

type Cidade struct {
	Id			int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	EstadoId	int64
	Nome		string
	Estado 		Estado 	`gorm:"ForeignKey:EstadoId; AssociationForeignKey:EstadoId"`
}

type Cidades []Cidade