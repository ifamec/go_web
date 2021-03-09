package utils

type UpdateCartResponse struct {
	TotalCount       int     `json:"total_count"`
	TotalAmount      float64 `json:"total_amount"`
	UpdateItemAmount float64 `json:"update_item_amount"`
}
