package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string               `bson:"name" json:"name"`
	Email        string               `bson:"email" json:"email"`
	GoogleID     string               `bson:"googleId" json:"-"`
	Role         string               `bson:"role" json:"role"`
	OrderHistory []primitive.ObjectID `bson:"orderHistory" json:"orderHistory"`
	CreatedAt    time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time            `bson:"updatedAt" json:"updatedAt"`
	LastLoginAt  time.Time            `bson:"lastLoginAt" json:"lastLoginAt"`
	Picture      string               `bson:"picture,omitempty" json:"picture,omitempty"`
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// NewUser creates a new user instance
func NewUser(name, email, googleID, role, picture string) *User {
	now := time.Now()
	return &User{
		Name:         name,
		Email:        email,
		GoogleID:     googleID,
		Role:         role,
		OrderHistory: []primitive.ObjectID{},
		CreatedAt:    now,
		UpdatedAt:    now,
		LastLoginAt:  now,
	}
}

// UpdateLastLogin updates the last login time of the user
func (u *User) UpdateLastLogin() {
	u.LastLoginAt = time.Now()
}

// AddOrder adds an order to the user's order history
func (u *User) AddOrder(orderID primitive.ObjectID) {
	u.OrderHistory = append(u.OrderHistory, orderID)
	u.UpdatedAt = time.Now()
}

// RemoveOrder removes an order from the user's order history
func (u *User) RemoveOrder(orderID primitive.ObjectID) {
	for i, id := range u.OrderHistory {
		if id == orderID {
			u.OrderHistory = append(u.OrderHistory[:i], u.OrderHistory[i+1:]...)
			break
		}
	}
	u.UpdatedAt = time.Now()
}

// IsAdmin checks if the user is an admin
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}
