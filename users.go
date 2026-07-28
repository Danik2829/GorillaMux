package main

import (
	"errors"
)

var (
	InvalidData = errors.New("invalid data")
	UserExist   = errors.New("user already exists")
	NotFound    = errors.New("User with this ID not exist")
)

type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
}

type UserService struct {
	Users map[string]User
}

func NewUserService() *UserService {
	return &UserService{
		Users: make(map[string]User),
	}
}

func (s *UserService) CreateUser(user User) error {
	if user.ID == "" || user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Age == 0 {
		return InvalidData
	}
	if _, exist := s.Users[user.ID]; exist {
		return UserExist
	}
	s.Users[user.ID] = user
	return nil
}

func (s *UserService) GetUser(id string) (User, error) {
	if _, exist := s.Users[id]; exist {
		return s.Users[id], nil
	}
	return User{}, NotFound
}

func (s *UserService) UpdateUser(user User) error {
	if user.ID == "" || user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Age == 0 {
		return InvalidData
	}
	if _, err := s.Users[user.ID]; !err {
		return NotFound
	} else {
		s.Users[user.ID] = user
	}
	return nil
}

func (s *UserService) DeleteUser(id string) error {
	if _, err := s.Users[id]; !err {
		return NotFound
	} else {
		delete(s.Users, id)
	}
	return nil
}

func (s *UserService) ListUsers() []User {
	users := make([]User, 0)
	for val := range s.Users {
		users = append(users, s.Users[val])
	}
	return users
}
