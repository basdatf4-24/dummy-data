package generatedata

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

type Contributor struct {
	Id              uuid.UUID
	Nama            string
	JenisKelamin    int
	Kewarganegaraan string
}

func GenerateContributors() (Contributor, error) {
	id := uuid.New()
	person := gofakeit.Person()
	country := gofakeit.Country()
	var genderNum int
	if person.Gender == "female" {
		genderNum = 1
	} else {
		genderNum = 0
	}
	name := person.FirstName + " " + person.LastName

	var contributor = Contributor{
		Id:              id,
		Nama:            name,
		JenisKelamin:    genderNum,
		Kewarganegaraan: country,
	}

	return contributor, nil
}
