package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus string

const (
	OrderStatusStarted   OrderStatus = "Started"
	OrderStatusCompleted OrderStatus = "Completed"
)

type Item struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name" binding:"required"`
	ImageURL string             `bson:"imageUrl" json:"imageUrl" binding:"required"`
}

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedBy primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	SendTo    primitive.ObjectID `bson:"sendTo" json:"sendTo"`
	Status    OrderStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Items     []Item             `bson:"items" json:"items" binding:"required"`
	OrderDate time.Time          `bson:"orderDate" json:"orderDate"`
}
