package generatedata

import "github.com/brianvoe/gofakeit/v7"

type User struct {
	Username string
	Password string
}

func GenerateUser() (User, error) {
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, false, false, 10)
	var user = User{
		Username: username,
		Password: password,
	}
	return user, nil
}
