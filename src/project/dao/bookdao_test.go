package dao

import (
	_ "database/sql"
	"fmt"
	"testing"
)

func TestBook (t *testing.T) {
	fmt.Println("Test - bookdao.go")
	t.Run("Get Books", testGetBooks)
}

func testGetBooks(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for k, v := range books {
		fmt.Println(k, v)
	}
}

