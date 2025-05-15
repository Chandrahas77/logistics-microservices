package warehouse

import (
	"context"
	"net/http"

	"github.com/Chandrahas77/logistics-microservices/warehouse-service/warehousepb"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, client warehousepb.WarehouseServiceClient) {
	rg.GET("/", func(c *gin.Context) {
		res, err := client.ListInventory(context.Background(), &warehousepb.ListInventoryRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res.Inventory)
	})

	rg.POST("/", func(c *gin.Context) {
		var req warehousepb.AddInventoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := client.AddInventory(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	rg.GET("/:item_id", func(c *gin.Context) {
		itemID := c.Param("item_id")
		res, err := client.GetInventory(context.Background(), &warehousepb.GetInventoryRequest{ItemId: itemID})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})
}
