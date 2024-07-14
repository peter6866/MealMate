package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpiceLevel string
type AlcoholContent string

const (
	SpiceLevelNone   SpiceLevel = "Not Spicy"
	SpiceLevelMild   SpiceLevel = "Mild"
	SpiceLevelMedium SpiceLevel = "Spicy"
	SpiceLevelHot    SpiceLevel = "Very Spicy"

	AlcoholContentNone AlcoholContent = "Non-Alcoholic"
	AlcoholContentHas  AlcoholContent = "Alcoholic"
)

type MenuItem struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string             `bson:"name" json:"name" binding:"required"`
	CategoryID     primitive.ObjectID `bson:"categoryId" json:"categoryId"`
	ImageURL       string             `bson:"imageUrl" json:"imageUrl"`
	SpiceLevel     SpiceLevel         `bson:"spiceLevel,omitempry" json:"spiceLevel,omitempty"`
	AlcoholContent AlcoholContent     `bson:"alcoholContent,omitempty" json:"alcoholContent,omitempty"`
	ReferenceLink  string             `bson:"referenceLink,omitempty" json:"referenceLink,omitempty"`
	CreatedBy      primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func (m *MenuItem) SetSpiceLevel(spiceLevel SpiceLevel) {
	m.SpiceLevel = spiceLevel
	m.UpdatedAt = time.Now()
}

func (m *MenuItem) SetAlcoholContent(alcoholContent AlcoholContent) {
	m.AlcoholContent = alcoholContent
	m.UpdatedAt = time.Now()
}

func (m *MenuItem) SetReferenceLink(referenceLink string) {
	m.ReferenceLink = referenceLink
	m.UpdatedAt = time.Now()
}
