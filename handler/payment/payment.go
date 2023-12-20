package payment

import (
	"fmt"

	"github.com/lenny-mo/router/global"
	"github.com/lenny-mo/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/lenny-mo/payment-api/proto/paymentapi"
)

func GetPaymentHandler(c *gin.Context) {

	// 开启链路追踪
	ctx, _ := middleware.ContextWithSpan(c)
	// 从body中获取参数
	var params GetPaymentParam
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("获取参数成功：", params.PaymentId)
	}

	client := paymentapi.NewPaymentAPIService("go.micro.api.payment-api", global.GetGlobalRPCService().Client())
	data, err := client.GetPayment(ctx, &paymentapi.GetPaymentRequest{
		PaymentId: params.PaymentId,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	c.String(200, data.PaymentInfo)
}

func CreatePaymentHandler(c *gin.Context) {
	// 开启链路追踪
	ctx, _ := middleware.ContextWithSpan(c)

	// 从body 中获取参数
	params := new(CreatePaymentParam)
	if err := c.ShouldBindJSON(params); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("获取参数成功：", params.Method, params.OrderId)
	}

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

func UpdatePaymentHandler(c *gin.Context) {
	// 开启链路追踪
	ctx, _ := middleware.ContextWithSpan(c)

	// 1. 获取参数
	var params UpdatePaymentParam
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("获取参数成功：", params.PaymentId, params.PaymentMethod, params.PaymenStatus)
	}

	client := paymentapi.NewPaymentAPIService("go.micro.api.payment-api", global.GetGlobalRPCService().Client())
	data, err := client.UpdatePayment(ctx, &paymentapi.UpdatePaymentRequest{
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
