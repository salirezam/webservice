package models

import "fmt"

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

func AddUser(user User) (User, error) {
	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}

func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, candidateUser := range users {
		if candidateUser.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", user.ID)
}

func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
