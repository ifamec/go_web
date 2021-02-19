package model

import (
	"fmt"
	"testing"
)

// TestMain actions before all tests
func TestMain(m *testing.M)  {
	fmt.Println("Before All Tests")

	// run Tests
	m.Run()
}

func TestUser(t *testing.T)  {
	fmt.Println("Start User model Test")

	// run child test
	// t.Run("Test Add User", testAddUser)
	t.Run("Test Get User By Id", testGetUserById)
	t.Run("Test Get Users", testGetUsers)
}

func testAddUser(t *testing.T) {
	// fmt.Println("Test AddUser")
	user := &User{}

	_ = user.AddUser()
	_ = user.AddUserNoPre()
}

func testGetUserById(t *testing.T) {
	user := &User{
		Id: 1,
	}

	user, _ = user.GetUserById()
	fmt.Println(user)
}

func testGetUsers(t *testing.T) {
	user := &User{}

	users, _ := user.GetUsers()
	for _, v := range users {
		fmt.Println(v)
	}
}