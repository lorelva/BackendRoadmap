package main

type UserReader interface {
	GetUserByID(id int) User
}

type UserWriter interface {
	CreateUser(user User) error
}

type UserDeleter interface {
	DeleteUserByID(id int) error
}

func main() {

}
