package dao

import (
	_ "database/sql"
	"fmt"
	"go_web/src/project/model"
	"testing"
)

func TestCartItem(t *testing.T) {
	fmt.Println("Test - cartitemdao.go")
	t.Run("Get Cart Item By BookId CartId", testGetCartItemByBookId)
	t.Run("Get Cart Items By CartId", testGetCartItemsByCartId)
	t.Run("Update Book Count", testUpdateBookCount)
}

func testGetCartItemByBookId(t *testing.T)  {
	cartItem, err := GetCartItemByBookIdCartId(1, "1234")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(cartItem)
}
func testUpdateBookCount(t *testing.T)  {
	book, _ := GetBookById(2)
	err := UpdateBookCount(&model.CartItem{
		Count: 7,
		Book: book,
	},1, "1234")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func testGetCartItemsByCartId(t *testing.T)  {
	cartItems, err := GetCartItemsByCartId("1234")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for _, v := range cartItems {
		fmt.Println(v)
	}
}

