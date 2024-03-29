package controller

import (
	Service "DailyFresh-Backend/domain/service/order"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/orders", Service.GetOrder)
	router.POST("/orders", Service.PostOrders)

	router.GET("/shipment", Service.GetShipments)
	router.PUT("/shipment/:order_id", Service.UpdateShipmentStatus)

	router.GET("/payment", Service.GetPayments)
}
