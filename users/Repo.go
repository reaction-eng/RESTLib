// Copyright 2019 Reaction Engineering International. All rights reserved.
// Use of this source code is governed by the MIT license in the file LICENSE.txt.

package users

//go:generate mockgen -destination=../mocks/mock_users_repo.go -package=mocks -mock_names Repo=MockUserRepo github.com/reaction-eng/restlib/users  Repo

type Repo interface {
	/**
	Get the user with the email.  An error is thrown is not found
	*/
	GetUserByEmail(email string) (User, error)

	/**
	Get the user with the ID.  An error is thrown is not found
	*/
	GetUser(id int) (User, error)

	/**
	Add User
	*/
	AddUser(user User) (User, error)

	/**
	Update User
	*/
	UpdateUser(user User) (User, error)

	/**
	Activate User
	*/
	ActivateUser(user User) error

	/**
	Create empty user
	*/
	NewEmptyUser() User

	/**
	List all users
	*/
	ListAllUsers() ([]int, error)
	ListAllActiveUsers() ([]int, error)
}
