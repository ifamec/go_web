package model

import "time"

type Order struct {
	OrderId     string
	Timestamp   time.Time
	TotalCount  int
	TotalAmount float64
	Status      int         // 0 order places, 1 shipped, 2 order complete
	UserId      int
}
