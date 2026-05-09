package domain

type UserRepository interface {
	Signup(user *User) error
	GetAll() ([]User, error)
	Login(user *User) error
}
