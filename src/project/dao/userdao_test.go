package dao

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestUser (t *testing.T) {
	fmt.Println("Test - userdao.go")
	t.Run("Login Validation", testLoginValidate)
	t.Run("Username Availability True", testUsernameAvailabilityTrue)
	t.Run("Username Availability False", testUsernameAvailabilityFalse)
	t.Run("Create User", testCreateUser)
}

func testLoginValidate(t *testing.T) {
	user, err := LoginValidate("admin", "12345678")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println("User Info:", user)
	}
}

func testUsernameAvailabilityTrue(t *testing.T) {
	isAvailable, err := UsernameAvailability("random_username")
	if err != nil && err == sql.ErrNoRows {
		fmt.Println("Username 'random_username' availability:", isAvailable)
	} else {
		fmt.Println(err)
		t.Fail()
	}
}

func testUsernameAvailabilityFalse(t *testing.T) {
	isAvailable, err := UsernameAvailability("admin")
	if err == nil {
		fmt.Println("Username 'random_username' availability:", isAvailable)
	} else {
		fmt.Println(err)
		t.Fail()
	}
}

func testCreateUser(t *testing.T) {
	err := CreateUser("test_user_01", "test_password", "test@test.com")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println("test_user_01 created")
	}
}
