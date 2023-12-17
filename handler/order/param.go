package order

type CreateOrderParam struct {
	OrderId      string `json:"order_id"`
	OrderVersion int64  `json:"order_version"`
	UserId       int64  `json:"user_id"`
	Status       int8   `json:"status"`
	OrderData    string `json:"order_data"`
}

type GetOrderParam struct {
	OrderId string `json:"orderid" binding:"required"`
}

type UpdateOrderParam struct {
	OldVersion   int64  `json:"old_version" binding:"required"`
	OrderVersion int64  `json:"order_version" binding:"required"`
	UserId       int64  `json:"user_id" binding:"required"`
	Status       int8   `json:"status"`
	OrderData    string `json:"order_data"`
	OrderId      string `json:"order_id" binding:"required"`
}
