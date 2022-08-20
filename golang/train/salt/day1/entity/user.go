package entity

import "allways-learning/golang/train/salt/day1/valueobject"

type User struct {
	Name     *valueobject.Name
	Email    *valueobject.Email
	Password *valueobject.Password
}

func NewUser(name *valueobject.Name, email *valueobject.Email, password *valueobject.Password) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
