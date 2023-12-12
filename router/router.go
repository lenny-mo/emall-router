package router

import (
	"github.com/lenny-mo/router/handler/cart"
	"github.com/lenny-mo/router/handler/order"
	"github.com/lenny-mo/router/handler/payment"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()

	// payment router group
	// // 暴露三个接口：结算订单，查询支付状态，更新支付状况
	paymentGroup := ginRouter.Group("/payment")
	{
		paymentGroup.POST("/create", payment.CreatePaymentHandler)
		paymentGroup.GET("/get", payment.GetPaymentHandler)
		paymentGroup.POST("update", payment.UpdatePaymentHandler)
	}

	// order router group
	// 暴露三个接口：创建订单，查询订单，更新订单
	orderGroup := ginRouter.Group("/order")
	{
		orderGroup.POST("/create", order.CreateOrderHandler)
		orderGroup.GET("/get", order.GetOrderHandler)
		orderGroup.POST("update", order.UpdateOrderHandler)
	}

	// cart router group
	// 暴露三个接口：加购，查询购物车
	cartGroup := ginRouter.Group("/cart")
	{
		cartGroup.POST("/add", cart.CreateCartHandler)
		cartGroup.GET("/get", cart.GetCartHandler)
	}

	return ginRouter
}
