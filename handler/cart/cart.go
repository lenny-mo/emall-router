package cart

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lenny-mo/cart-api/proto/cartapi"
	"github.com/lenny-mo/router/global"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateCartParam struct {
	UserId   string    `json:"userId" binding:"required"`
	SKUId    string    `json:"skuId" binding:"required"`
	Time     time.Time `json:"time,omitempty" binding:"required"`
	Status   int8      `json:"status"`
	Quantity int32     `json:"quantity" binding:"required"`
}

func CreateCartHandler(c *gin.Context) {
	// 1. 获取参数
	var params CreateCartParam
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		return
	}

	// 2. 构造client
	client := cartapi.NewCartApiService("go.micro.api.cart-api", global.GetGlobalRPCService().Client())
	data, err := client.Add(context.TODO(), &cartapi.AddCartRequest{
		UserId: params.UserId,
		Item: &cartapi.CartItem{
			Skuid:    params.SKUId,
			Quantity: params.Quantity,
			Time:     timestamppb.New(params.Time),
			Status:   cartapi.CartStatus(params.Status),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

type GetCartParam struct {
	UserId string `json:"user_id" binding:"required"`
}

func GetCartHandler(c *gin.Context) {
	// 1. 获取参数
	var param GetCartParam
	if err := c.ShouldBindJSON(&param); err != nil {
		fmt.Println(err)
		return
	}

	// 2. 构造client
	client := cartapi.NewCartApiService("go.micro.api.cart-api", global.GetGlobalRPCService().Client())
	data, err := client.FindAll(context.TODO(), &cartapi.FindAllRequest{
		Userid: param.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

type UpdateCartParams struct {
	UserId   string    `json:"userId" binding:"required"`
	SKUId    string    `json:"skuId" binding:"required"`
	Time     time.Time `json:"time,omitempty" binding:"required"`
	Status   int8      `json:"status"`
	Quantity int32     `json:"quantity" binding:"required"`
}

func UpdateCartHandler(c *gin.Context) {
	// 1. 获取参数
	var params UpdateCartParams
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		return
	}

	// 2. 构造client
	client := cartapi.NewCartApiService("go.micro.api.cart-api", global.GetGlobalRPCService().Client())
	data, err := client.Update(context.TODO(), &cartapi.UpdateRequest{
		UserId: params.UserId,
		Item: &cartapi.CartItem{
			Skuid:    params.SKUId,
			Quantity: params.Quantity,
			Time:     timestamppb.New(params.Time),
			Status:   cartapi.CartStatus(params.Status),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

type DeleteCartParam struct {
	UserId string `json:"userId" binding:"required"`
	SKUId  string `json:"skuId" binding:"required"`
}

func DeleteCartHandler(c *gin.Context) {
	// 1. 获取参数
	var params DeleteCartParam
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Println(err)
		return
	}
	// 2. 构造client
	client := cartapi.NewCartApiService("go.micro.api.cart-api", global.GetGlobalRPCService().Client())
	data, err := client.Delete(context.TODO(), &cartapi.DeleteRequest{
		Userid: params.UserId,
		Skuid:  params.SKUId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
