package dao

import (
	_ "database/sql"
	"fmt"
	"go_web/src/project/model"
	"testing"
)

func TestCart(t *testing.T) {
	fmt.Println("Test - cartdao.go, cartitemsdao.go")
	t.Run("Add Cart", testAddCart)
}

// 1,1984 (Signet Classics),George Orwell,9.99,100,100,41aM4xOZxaL
// 2,A Brief History of Time,Stephen Hawking,18,100,100,51+GySc8ExL
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
