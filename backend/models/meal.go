package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MealTpye string

const (
	Breakfast MealTpye = "Breakfast"
	Lunch     MealTpye = "Lunch"
	Dinner    MealTpye = "Dinner"
	Snakcs    MealTpye = "Snacks"
)

type Meal struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MealDate    primitive.DateTime `bson:"mealDate" json:"mealDate" binding:"required"`
	MealType    MealTpye           `bson:"mealType" json:"mealType" binding:"required"`
	PhotoURL    string             `bson:"photoURL" json:"photoURL"`
	Items       []Item             `bson:"items" json:"items" binding:"required"`
	CreatedBy   primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	WithPartner bool               `bson:"withPartner" json:"withPartner" binding:"required"`
}
