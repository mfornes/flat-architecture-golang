package main

type Storage interface {
	SaveUser(user User) error
	FindUser(id int) (User, error)
	FindUsers() ([]User, error)
	DeleteUser(id int) error
}

func NewStorage() Storage {
	return new(StorageMemory)
}
