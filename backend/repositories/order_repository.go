package repositories

import (
	"context"

	"github.com/peter6866/foodie/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(client *mongo.Client) *OrderRepository {
	collection := client.Database("foodie").Collection("orders")
	return &OrderRepository{collection: collection}
}

// create a new order
func (r *OrderRepository) Create(ctx context.Context, order *models.Order) error {
	result, err := r.collection.InsertOne(ctx, order)
	if err == nil {
		order.ID = result.InsertedID.(primitive.ObjectID)
	}

	return err
}

// get all orders with createdBy for non-chef user
func (r *OrderRepository) GetAllForUser(ctx context.Context, createdBy primitive.ObjectID) ([]*models.Order, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"createdBy": createdBy})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*models.Order
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// get all orders with sendTo for chef user
func (r *OrderRepository) GetAllForChef(ctx context.Context, sendTo primitive.ObjectID) ([]*models.Order, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"sendTo": sendTo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*models.Order
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
