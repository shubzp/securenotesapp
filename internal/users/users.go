package users

import (
	"fmt"

	"github.com/hashicorp/go-uuid"
)

type User struct {
	UserId   string
	Name     string
	PhoneNo  string
	UserName string
	Password string
}

func MakeNewUser(name, phoneno, username, password string) IUser {
	uuid, _ := uuid.GenerateUUID()
	user := &User{}
	user.UserId = uuid
	user.Name = name
	user.PhoneNo = phoneno
	user.UserName = username
	user.Password = password
	return user
}

func (user *User) GetUser() User {
	return *user
}

func (user User) ShowUser() {
	fmt.Printf("Name : %s; PhoneNo : %s", user.Name, user.PhoneNo)
}
