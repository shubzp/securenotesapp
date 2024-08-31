package users

import "fmt"

type User struct {
	Name    string
	PhoneNo string
}

func MakeNewUser() IUser {
	return &User{}
}

func (user *User) Login(username, password string) bool {
	return true
}

func (user *User) Register(name, phoneno, username, password string) bool {
	return true
}

func (user User) ShowUser() {
	fmt.Printf("Name : %s; PhoneNo : %s", user.Name, user.PhoneNo)
}
