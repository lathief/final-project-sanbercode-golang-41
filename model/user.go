package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID         int
	Username   string
	Password   string
	CustomerID int
}

type UserCustomer struct {
	ID           int
	CustomerID   int
	Username     string
	Email        string
	MobileNumber string
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
