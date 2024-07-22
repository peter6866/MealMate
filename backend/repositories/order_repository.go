package repositories

import (
	"context"

	"github.com/peter6866/foodie/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// get order by ID
func (r *OrderRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Order, error) {
	var order models.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// get all orders with createdBy for non-chef user
func (r *OrderRepository) GetAllForUser(ctx context.Context, createdBy primitive.ObjectID) ([]*models.Order, error) {
	opts := options.Find().SetSort(bson.D{{Key: "orderDate", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{"createdBy": createdBy}, opts)
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
	opts := options.Find().SetSort(bson.D{{Key: "orderDate", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{"sendTo": sendTo}, opts)
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

// Complete an order sendTo the userID of the chef
func (r *OrderRepository) CompleteOrder(ctx context.Context, orderID primitive.ObjectID, userID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": orderID, "sendTo": userID}, bson.M{"$set": bson.M{"status": models.OrderStatusCompleted}})
	return err
}
