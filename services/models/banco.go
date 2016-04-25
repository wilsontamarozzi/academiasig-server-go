package models

type Banco struct {
	Id 		int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	Numero	string
	Nome	string
}

type Bancos []Banco

func (b Banco) IsValid() map[string][]string {

	err := make(map[string][]string)

	if b.Nome == "" {
		err["nome"] = append(err["nome"], "Nome não pode ser vázio.")
	}

	if len(b.Nome) < 2 {
		err["nome"] = append(err["nome"], "Nome deve ter mais de 2 caracteres.")
	}

	if b.Numero == "" {
		err["numero"] = append(err["numero"], "Numero não pode ser vázio.")	
	}

	return err
}