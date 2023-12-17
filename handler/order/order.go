package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenny-mo/order-api/proto/orderapi"
	"github.com/lenny-mo/router/global"
)

func CreateOrderHandler(c *gin.Context) {
	// 1. 解析参数
	params := new(CreateOrderParam)
	if err := c.ShouldBindJSON(params); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("获取参数成功：", params.OrderData, params.OrderId, params.OrderVersion, params.Status, params.UserId)
	}

	// 2. 构造orderapi的客户端
	client := orderapi.NewOrderApiService("go.micro.api.order-api", global.GetGlobalRPCService().Client())
	data, err := client.CreateOrder(context.TODO(), &orderapi.CreateOrderRequest{
		Data: &orderapi.OrderInfo{
			OrderId:      params.OrderId,
			OrderVersion: params.OrderVersion,
			UserId:       params.UserId,
			Status:       orderapi.OrderStatus(params.Status),
			OrderData:    params.OrderData,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data.Rowaffectd: ", data.Rowaffectd)
}

func GetOrderHandler(c *gin.Context) {
	// 1. 解析参数
	var params GetOrderParam
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		return
	}

	// 2. 请求order-api的get方法
	client := orderapi.NewOrderApiService("go.micro.api.order-api", global.GetGlobalRPCService().Client())
	data, err := client.GetOrder(context.TODO(), &orderapi.GetOrderRequest{
		Orderid: params.OrderId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data: ", data)
}

func UpdateOrderHandler(c *gin.Context) {
	// 1. 获取参数
	var params UpdateOrderParam
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		return
	}

	client := orderapi.NewOrderApiService("go.micro.api.order-api", global.GetGlobalRPCService().Client())
	data, err := client.UpdateOrder(context.TODO(), &orderapi.UpdateOrderRequest{
		Data: &orderapi.OrderInfo{
			OrderId:      params.OrderId,
			OrderVersion: params.OrderVersion,
			UserId:       params.UserId,
			Status:       orderapi.OrderStatus(params.Status),
			OrderData:    params.OrderData,
		},
		Oldversion: params.OldVersion,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data.Rowaffectd: ", data.Rowaffectd)
}

func GetUUIDHandler(c *gin.Context) {
	client := orderapi.NewOrderApiService("go.micro.api.order-api", global.GetGlobalRPCService().Client())
	data, err := client.GenerateUUID(context.TODO(), &orderapi.Empty{})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, "orderId: "+string(data.Uuid))
	fmt.Println(data)
}
