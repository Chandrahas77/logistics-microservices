package shipment

import (
	"context"
	"net/http"

	"github.com/Chandrahas77/logistics-microservices/shipment-service/shipmentpb"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, client shipmentpb.ShipmentServiceClient) {
	rg.GET("/", func(c *gin.Context) {
		res, err := client.ListShipments(context.Background(), &shipmentpb.ListShipmentsRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res.Shipments)
	})

	rg.GET("/:id", func(c *gin.Context) {
		shipmentID := c.Param("id")
		res, err := client.GetShipment(context.Background(), &shipmentpb.GetShipmentRequest{ShipmentId: shipmentID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	rg.POST("/", func(c *gin.Context) {
		var req shipmentpb.CreateShipmentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := client.CreateShipment(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})
}
