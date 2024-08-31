package users

type IUser interface {
	Login(username string, password string) bool
	Register(name string, phoneno string, username string, password string) bool
}
