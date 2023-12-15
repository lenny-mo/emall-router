package payment

type GetPaymentParam struct {
	PaymentId string `json:"paymentId" binding:"required"`
}

type CreatePaymentParam struct {
	OrderId int64  `json:"orderId" binding:"required"`
	Method  string `json:"method" binding:"required"`
}

type UpdatePaymentParam struct {
	PaymentId     string `json:"paymentId" binding:"required"`
	PaymentMethod string `json:"method"`
	PaymenStatus  int32  `json:"status"`
}
