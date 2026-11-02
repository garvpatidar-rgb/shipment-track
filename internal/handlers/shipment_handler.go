package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shiptrack/internal/models"
)

type ShipmentHandler struct {
	shipments []models.Shipment
	nextID    uint
}

func NewShipmentHandler() *ShipmentHandler {
	return &ShipmentHandler{
		shipments: []models.Shipment{},
		nextID:    1,
	}
}

func (h *ShipmentHandler) CreateShipment(c *gin.Context) {
	var shipment models.Shipment

	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shipment.ID = h.nextID
	h.nextID++

	if shipment.Status == "" {
		shipment.Status = "pending"
	}

	h.shipments = append(h.shipments, shipment)

	c.JSON(http.StatusCreated, shipment)
}

func (h *ShipmentHandler) GetAllShipments(c *gin.Context) {
	c.JSON(http.StatusOK, h.shipments)
}
