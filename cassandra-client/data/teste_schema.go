package data

import (
	"testing"
)

func BenchmarkMakeAndIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := make([]User, 1000)
		var user User
		for idx := 0; idx < 1000; idx++ {
			users[idx] = user
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var users []User
		var user User
		for idx := 0; idx < 1000; idx++ {
			users = append(users, user)
		}
	}
}
