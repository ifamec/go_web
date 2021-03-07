package model

type CartItem struct {
	CartItemId int
	Book       *Book
	Count      int
	Amount     float64 // Count * Book.Price
	CartId     string  // uuid
}

func (c *CartItem) GetAmount() (amount float64) {
	return c.Book.Price * float64(c.Count)
}