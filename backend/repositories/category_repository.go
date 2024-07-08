package repositories

import (
	"context"

	"github.com/peter6866/foodie/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(client *mongo.Client) *CategoryRepository {
	collection := client.Database("foodie").Collection("categories")
	return &CategoryRepository{collection: collection}
}

func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	result, err := r.collection.InsertOne(ctx, category)
	if err == nil {
		category.ID = result.InsertedID.(primitive.ObjectID)
	}
	return err
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]*models.Category, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []*models.Category
	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error) {
	var category models.Category
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
