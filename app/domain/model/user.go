package model

import "github.com/yuuuutsk/gobase-backend/app/domain"

type User struct {
	ID        domain.UserID
	FirstName string
	LastName  string
}

func NewUser(
	FirstName string,
	LastName string,
) *User {
	return &User{
		FirstName: FirstName,
		LastName:  LastName,
	}
}

func RestoreUser(
	ID domain.UserID,
	FirstName string,
	LastName string,
) *User {
	return &User{
		ID:        ID,
		FirstName: FirstName,
		LastName:  LastName,
	}
}
