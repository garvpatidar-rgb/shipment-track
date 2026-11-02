package models

type Shipment struct {
	ID           uint   `json:"id"`
	SenderName   string `json:"sender_name" binding:"required"`
	ReceiverName string `json:"receiver_name" binding:"required"`
	Origin       string `json:"origin" binding:"required"`
	Destination  string `json:"destination" binding:"required"`
	Status       string `json:"status"`
}