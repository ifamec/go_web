package dao

import (
	_ "database/sql"
	"fmt"
	"go_web/src/project/model"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	fmt.Println("Test - orderdao.go")
	t.Run("Add Order", testAddOrder)
	t.Run("Get My Orders", testGetMyOrders)
	t.Run("update Order State", testUpdateOrderState)
}

func testAddOrder(t *testing.T) {
	const OrderId = "1234"
	oi1 := &model.OrderItem{
		Count: 1,
		Amount: 300,
		Title: "TestTitle_1",
		Author: "TestAuthor_1",
		Price: 300,
		ImgPath: "TestImgPath_1",
		OrderId: OrderId,
	}
	oi2 := &model.OrderItem{
		Count: 1,
		Amount: 100,
		Title: "TestTitle_2",
		Author: "TestAuthor_2",
		Price: 100,
		ImgPath: "TestImgPath_2",
		OrderId: OrderId,
	}
	err := AddOrder(&model.Order{
		OrderId:     OrderId,
		Timestamp:   time.Now(),
		TotalCount:  2,
		TotalAmount: 400,
		Status:      0,
		UserId:      1,
		Username:    "admin",
	})
	err = AddOrderItem(oi1)
	err = AddOrderItem(oi2)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func testGetOrders(t *testing.T) {
	orders, err := GetOrders()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}
func testGetMyOrders(t *testing.T) {
	orders, err := GetMyOrders(1, "admin")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}
func testUpdateOrderState(t *testing.T) {
	err := UpdateOrderState("1234", 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}