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
