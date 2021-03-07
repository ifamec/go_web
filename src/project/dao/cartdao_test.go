package dao

import (
	_ "database/sql"
	"fmt"
	"go_web/src/project/model"
	"testing"
)

func TestCart(t *testing.T) {
	fmt.Println("Test - cartdao.go")
	t.Run("Add Cart", testAddCart)
}

func testAddCart(t *testing.T) {
	book1 := &model.Book{Id: 1, Price: 9.99}
	book2 := &model.Book{Id: 2, Price: 18.00}

	cart1 := &model.CartItem{Book: book1, Count: 2, CartId: "1234"}
	cart2 := &model.CartItem{Book: book2, Count: 5, CartId: "1234"}

	cartItems := make([]*model.CartItem, 0)
	cartItems = append(cartItems, cart1)
	cartItems = append(cartItems, cart2)

	cart := &model.Cart{
		CartId:    "1234",
		CartItems: cartItems,
		UserId:    2,
	}

	err := AddCart(cart)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
