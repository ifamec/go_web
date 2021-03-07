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
	t.Run("Get Session", testGetSession)
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

func testGetSession(t *testing.T) {
	session, err := GetSession("6ee684e3-99df-4ef0-6409-f36f1e5b4b1c")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(session.SessionId, session.Username, session.UserId)
}
func testDeleteSession(t *testing.T) {
	err := DeleteSession("test-session-id")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}