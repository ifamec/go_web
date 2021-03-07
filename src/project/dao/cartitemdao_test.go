package dao

import (
	_ "database/sql"
	"fmt"
	"testing"
)

func TestCartItem(t *testing.T) {
	fmt.Println("Test - cartitemdao.go")
	t.Run("Get Cart Item By BookId", testGetCartItemByBookId)
	t.Run("Get Cart Items By CartId", testGetCartItemsByCartId)
}

func testGetCartItemByBookId(t *testing.T)  {
	cartItem, err := GetCartItemByBookId(2, "1234")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(cartItem)
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

