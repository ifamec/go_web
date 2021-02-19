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
	t.Run("Test Add User", testAddUser)
}

func testAddUser(t *testing.T) {
	// fmt.Println("Test AddUser")
	user := &User{}

	_ = user.AddUser()
	_ = user.AddUserNoPre()
}