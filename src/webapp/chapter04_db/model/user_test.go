package model

import (
	"fmt"
	"testing"
)

func TestUser_AddUser(t *testing.T) {
	fmt.Println("Test AddUser")
	user := &User{}

	_ = user.AddUser()
	_ = user.AddUserNoPre()
}