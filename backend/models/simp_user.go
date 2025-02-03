package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SimpUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email        string             `bson:"email" json:"email"`
	IsChef       bool               `bson:"isChef,omitempty" json:"isChef,omitempty"`
	PartnerEmail string             `bson:"partnerEmail,omitempty" json:"partnerEmail,omitempty"`
}
