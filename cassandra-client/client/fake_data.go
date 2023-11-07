package client

import (
	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Name   string
	Gender string
	Phone  string
}

func GetFakeUsers(number int) []User {
	users := make([]User, number)

	for i := 0; i < number; i++ {
		users[i].Name = gofakeit.FirstName()
		users[i].Gender = gofakeit.Gender()
		users[i].Phone = gofakeit.Phone()
	}

	return users
}
