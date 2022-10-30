package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index:email,unique"`
	Password string `gorm:"password"`
	Name     string `gorm:"name"`
}

type UserRepository interface {
	Create(email string, password string, name string) (*User, error)
	GetAll() ([]User, error)
	Login() (*User, error)
	GetUserByID(userID int) (*User, error)
	GetUserByEmail(email string) (*User, error)
}
