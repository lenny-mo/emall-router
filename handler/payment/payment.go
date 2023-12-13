package payment

import (
	"context"
	"fmt"
	"time"

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

type CreatePaymentParam struct {
	OrderId int64  `json:"orderId" binding:"required"`
	Method  string `json:"method" binding:"required"`
}

func CreatePaymentHandler(c *gin.Context) {
	// 从body 中获取参数
	params := new(CreatePaymentParam)
	if err := c.ShouldBindJSON(params); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("获取参数成功：", params.Method, params.OrderId)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client := paymentapi.NewPaymentAPIService("go.micro.api.payment-api", global.GetGlobalRPCService().Client())
	data, err := client.MakePayment(ctx, &paymentapi.MakePaymentRequest{
		Method:  params.Method,
		OrderId: params.OrderId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

type UpdatePaymentParam struct {
	PaymentId     string `json:"paymentId" binding:"required"`
	PaymentMethod string `json:"method"`
	PaymenStatus  int32  `json:"status"`
}

func UpdatePaymentHandler(c *gin.Context) {
	// 1. 获取参数
	var params UpdatePaymentParam
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("获取参数成功：", params.PaymentId, params.PaymentMethod, params.PaymenStatus)
	}

	client := paymentapi.NewPaymentAPIService("go.micro.api.payment-api", global.GetGlobalRPCService().Client())
	data, err := client.UpdatePayment(context.Background(), &paymentapi.UpdatePaymentRequest{
		PaymentId:     params.PaymentId,
		PaymentStatus: params.PaymenStatus,
		PaymentMethod: params.PaymentMethod,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
