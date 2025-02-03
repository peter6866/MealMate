package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string               `bson:"name" json:"name"`
	Email        string               `bson:"email" json:"email"`
	GoogleId     string               `bson:"googleId" json:"-"`
	Role         string               `bson:"role" json:"role"`
	OrderHistory []primitive.ObjectID `bson:"orderHistory" json:"orderHistory"`
	CreatedAt    time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time            `bson:"updatedAt" json:"updatedAt"`
	LastLoginAt  time.Time            `bson:"lastLoginAt" json:"lastLoginAt"`
	Picture      string               `bson:"picture,omitempty" json:"picture,omitempty"`
	IsChef       bool                 `bson:"isChef,omitempty" json:"isChef,omitempty"`
	PartnerEmail string               `bson:"partnerEmail,omitempty" json:"partnerEmail,omitempty"`
	Cart         []primitive.ObjectID `bson:"cart" json:"cart"`
}

const (
	RoleChef = "chef"
	RoleUser = "user"
)

// NewUser creates a new user instance
func NewUser(name, email, googleId, role, picture string) *User {
	now := time.Now()
	return &User{
		Name:         name,
		Email:        email,
		GoogleId:     googleId,
		Role:         role,
		OrderHistory: []primitive.ObjectID{},
		CreatedAt:    now,
		UpdatedAt:    now,
		LastLoginAt:  now,
		Picture:      picture,
		Cart:         []primitive.ObjectID{},
	}
}

// UpdateLastLogin updates the last login time of the user
func (u *User) UpdateLastLogin() {
	u.LastLoginAt = time.Now()
}

// IsAdmin checks if the user is an admin
func (u *User) IsAdmin() bool {
	return u.Role == RoleChef
}
