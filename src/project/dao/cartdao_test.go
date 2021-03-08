package dao

import (
	_ "database/sql"
	"fmt"
	"testing"
)

func TestCart(t *testing.T) {
	fmt.Println("Test - cartdao.go")
	// t.Run("Add Cart", testAddCart)
	t.Run("Get Cart By UserId", testGetCartByUserId)
	t.Run("Detele Cart By Cart", testDeleteCartByCartId)
}

func testGetCartByUserId(t *testing.T) {
	cart, err := GetCartByUserId(2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(cart.CartId, cart.TotalCount, cart.TotalAmount, cart.UserId)
	for _, v := range cart.CartItems {
		fmt.Println(v)
	}
}

func testDeleteCartByCartId(t *testing.T) {
	err := DeleteCartByCartId("b6891c8b-35c1-4e23-5e45-a8a82da711f8")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
