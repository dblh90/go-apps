package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {

	if u.ID != 0 {
		return User{}, errors.New("new user must not include id")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID %v not found", id)
}

func UpdateUser(userToUpdate User) (User, error) {

	for index, user := range users {
		if userToUpdate.ID == user.ID {
			users[index] = &userToUpdate
			return userToUpdate, nil
		}
	}
	return User{}, fmt.Errorf("User with ID %v not found", userToUpdate.ID)
}

func RemoveUserById(id int) error {
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
