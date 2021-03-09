package dao

import (
	_ "database/sql"
	"fmt"
	"testing"
)

func TestOrderItem(t *testing.T) {
	fmt.Println("Test - orderitemdao.go")
	t.Run("Get OrderItems By OrderId", testGetOrderItemsByOrderId)
}

func testGetOrderItemsByOrderId(t *testing.T) {
	orderItems, err := GetOrderItemsByOrderId("1234", "admin")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for _, v := range orderItems {
		fmt.Println(v)
	}
}