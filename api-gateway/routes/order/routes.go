package order

import (
	"context"
	"net/http"

	"github.com/Chandrahas77/logistics-microservices/order-service/orderpb"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, client orderpb.OrderServiceClient) {
	rg.GET("/", func(c *gin.Context) {
		res, err := client.ListOrders(context.Background(), &orderpb.ListOrdersRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res.Orders)
	})

	rg.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := client.GetOrder(context.Background(), &orderpb.GetOrderRequest{OrderId: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	rg.POST("/", func(c *gin.Context) {
		var req orderpb.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := client.CreateOrder(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})
}
