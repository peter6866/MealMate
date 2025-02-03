package repositories

import (
	"context"

	"github.com/peter6866/foodie/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SimpUserRepository struct {
	collection *mongo.Collection
}

func NewSimpUserRepository(client *mongo.Client) *SimpUserRepository {
	collection := client.Database("foodie").Collection("simp_users")
	return &SimpUserRepository{collection: collection}
}

func (r *SimpUserRepository) Create(ctx context.Context, user *models.SimpUser) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

// Upsert function using ObjectID as the unique identifier
func (r *SimpUserRepository) Upsert(ctx context.Context, user *models.SimpUser) error {
	filter := bson.M{"_id": user.ID} // Find document by ObjectID
	update := bson.M{"$set": user}   // Update fields
	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *SimpUserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.SimpUser, error) {
	var user models.SimpUser
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SimpUserRepository) FindByEmail(ctx context.Context, email string) (*models.SimpUser, error) {
	var user models.SimpUser
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
