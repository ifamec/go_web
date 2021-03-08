package model

type Cart struct {
	CartId      string
	CartItems   []*CartItem
	TotalCount  int
	TotalAmount float64
	UserId      int
	Username    string
}

func (c *Cart) GetTotalAmount() (tamount float64) {
	for _, v := range c.CartItems {
		tamount += v.GetAmount()
	}
	return
}
func (c *Cart) GetTotalCount() (tcount int) {
	for _, v := range c.CartItems {
		tcount += v.Count
	}
	return
}
