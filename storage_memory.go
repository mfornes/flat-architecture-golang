package main

import "time"

type StorageMemory struct {
	users []User
	count int
}

func (s *StorageMemory) SaveUser(user User) error {
	for _, u := range s.users {
		if u.CI == user.CI {
			return NewError("User already exists")
		}
	}
	user.ID = s.count + 1
	user.Created = time.Now()
	s.users = append(s.users, user)
	s.count++
	return nil
}

func (s *StorageMemory) DeleteUser(id int) error {
	for i, u := range s.users {
		if u.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return NewError("User not exists")
}

func (s *StorageMemory) FindUser(id int) (User, error) {
	var user User
	for _, u := range s.users {
		if u.ID == id {
			user = u
			return user, nil
		}
	}
	return user, NewError("User not exists")

}

func (s *StorageMemory) FindUsers() ([]User, error) {
	return s.users, nil
}
