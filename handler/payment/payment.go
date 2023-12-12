package payment

import (
	"context"
	"fmt"

	"github.com/lenny-mo/router/global"

	"github.com/gin-gonic/gin"
	"github.com/lenny-mo/payment-api/proto/paymentapi"
)

func GetPaymentHandler(c *gin.Context) {
	// 在这个函数内部使用payment api 的函数
	// 尝试使用rpc请求
	client := paymentapi.NewPaymentAPIService("go.micro.api.payment-api", global.GetGlobalRPCService().Client())
	data, err := client.GetPayment(context.Background(), &paymentapi.GetPaymentRequest{
		PaymentId: "12",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	c.String(200, data.PaymentInfo)
}

func CreatePaymentHandler(c *gin.Context) {
	client := paymentapi.NewPaymentAPIService("go.micro.api.payment-api", global.GetGlobalRPCService().Client())
	data, err := client.MakePayment(context.Background(), &paymentapi.MakePaymentRequest{
		Method:  "paypal",
		OrderId: 0,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func UpdatePaymentHandler(c *gin.Context) {
	client := paymentapi.NewPaymentAPIService("go.micro.api.payment-api", global.GetGlobalRPCService().Client())
	data, err := client.UpdatePayment(context.Background(), &paymentapi.UpdatePaymentRequest{
		PaymentId:     "",
		PaymentStatus: 1,
		PaymentMethod: "paypal",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
