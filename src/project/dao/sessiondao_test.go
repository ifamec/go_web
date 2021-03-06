package dao

import (
	_ "database/sql"
	"fmt"
	"go_web/src/project/model"
	"testing"
)

func TestSessions (t *testing.T) {
	fmt.Println("Test - sessiondao.go")
	t.Run("Add Session", testAddSession)
	t.Run("Delete Session", testDeleteSession)
}

func testAddSession(t *testing.T) {
	err := AddSession(&model.Session{
		SessionId: "test-session-id",
		Username: "test_user",
		UserId: 7,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func testDeleteSession(t *testing.T) {
	err := DeleteSession("test-session-id")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}