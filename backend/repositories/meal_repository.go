package repositories

import (
	"context"

	custom_errors "github.com/peter6866/foodie/custom-errors"
	"github.com/peter6866/foodie/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MealRepository struct {
	collection *mongo.Collection
}

func NewMealRepository(client *mongo.Client) *MealRepository {
	collection := client.Database("foodie").Collection("meals")
	return &MealRepository{collection: collection}
}

// log a new meal
func (r *MealRepository) Create(ctx context.Context, meal *models.Meal) error {
	result, err := r.collection.InsertOne(ctx, meal)
	if err == nil {
		meal.ID = result.InsertedID.(primitive.ObjectID)
	}

	return err
}

// update a meal
func (r *MealRepository) Update(ctx context.Context, meal *models.Meal) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": meal.ID}, meal)
	return err
}

// delete a meal
func (r *MealRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// find a meal by ID
func (r *MealRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Meal, error) {
	var meal models.Meal
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&meal)
	// if meal not found return not found error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, custom_errors.ErrMealNotFound
		}
		return nil, err
	}

	return &meal, nil
}

// find meals by date range and createdBy userID (e.g. 7/23 - 7/30)
func (r *MealRepository) FindByDateRange(ctx context.Context, startDate primitive.DateTime, endDate primitive.DateTime, createdBy primitive.ObjectID) ([]*models.Meal, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"mealDate": bson.M{"$gte": startDate, "$lte": endDate}, "createdBy": createdBy})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var meals []*models.Meal
	if err := cursor.All(ctx, &meals); err != nil {
		return nil, err
	}

	return meals, nil
}

// get all meals with createdBy userID order by mealDate desc
func (r *MealRepository) GetAllForUser(ctx context.Context, createdBy primitive.ObjectID) ([]*models.Meal, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"createdBy": createdBy}, options.Find().SetSort(bson.M{"mealDate": -1}))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var meals []*models.Meal
	if err := cursor.All(ctx, &meals); err != nil {
		return nil, err
	}

	return meals, nil
}
