package repositories

import (
	"context"

	"github.com/peter6866/foodie/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	collection := client.Database("foodie").Collection("users")
	return &UserRepository{collection: collection}
}

// create a new user
func (r *UserRepository) Create(user *models.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

// find a user by ID
func (r *UserRepository) FindByID(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// find by googleID
func (r *UserRepository) FindByGoogleID(googleId string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"googleId": googleId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update a user
func (r *UserRepository) Update(user *models.User) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	return err
}

// Add an order to a user's history
func (r *UserRepository) AddOrder(userID, orderID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$push": bson.M{"orderHistory": orderID}},
	)
	return err
}

// Remove an order from a user's history
func (r *UserRepository) RemoveOrder(userID, orderID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$pull": bson.M{"orderHistory": orderID}},
	)
	return err
}
