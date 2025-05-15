package main

import (
	"github.com/Chandrahas77/logistics-microservices/api-gateway/clients"
	orderRoutes "github.com/Chandrahas77/logistics-microservices/api-gateway/routes/order"
	shipmentRoutes "github.com/Chandrahas77/logistics-microservices/api-gateway/routes/shipment"
	warehouseRoutes "github.com/Chandrahas77/logistics-microservices/api-gateway/routes/warehouse"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	c := clients.InitClients()

	orderGroup := r.Group("/orders")
	orderRoutes.RegisterRoutes(orderGroup, c.OrderClient)

	shipmentGroup := r.Group("/shipments")
	shipmentRoutes.RegisterRoutes(shipmentGroup, c.ShipmentClient)

	warehouseGroup := r.Group("/inventory")
	warehouseRoutes.RegisterRoutes(warehouseGroup, c.WarehouseClient)

	r.Run(":8080")
}
