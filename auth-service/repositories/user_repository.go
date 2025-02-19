package repositories

import (
	"auth-service/config"
	"auth-service/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	collection := config.MongoDB.Collection("users")
	return &UserRepository{collection: collection}
}

// create a new user and return its id
func (r *UserRepository) Create(ctx context.Context, user *models.User) (primitive.ObjectID, error) {
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// find a user by ID
func (r *UserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// find by googleID
func (r *UserRepository) FindByGoogleID(ctx context.Context, googleId string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"googleId": googleId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// find by email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	// if user does not exist, return not found error
	if err == mongo.ErrNoDocuments {
		return nil, mongo.ErrNoDocuments
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update a user
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	return err
}

// Add a menu item to a user's cart
func (r *UserRepository) AddToCart(ctx context.Context, userID, menuItemID primitive.ObjectID) error {
	// Check if the item is already in the cart
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return err
	}
	for _, item := range user.Cart {
		if item == menuItemID {
			return nil
		}
	}

	_, err = r.collection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{"$push": bson.M{"cart": menuItemID}},
	)
	return err
}

// delete a menu item from a user's cart
func (r *UserRepository) RemoveFromCart(ctx context.Context, userID, menuItemID primitive.ObjectID) error {

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{"$pull": bson.M{"cart": menuItemID}},
	)

	return err
}

// clear the user's cart
func (r *UserRepository) ClearCart(ctx context.Context, userID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"cart": []primitive.ObjectID{}}},
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
