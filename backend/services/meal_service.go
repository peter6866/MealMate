package services

import (
	"context"
	"mime/multipart"

	custom_errors "github.com/peter6866/foodie/custom-errors"
	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	"github.com/peter6866/foodie/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MealService struct {
	userRepo *repositories.UserRepository
	mealRepo *repositories.MealRepository
}

func NewMealService(userRepo *repositories.UserRepository, mealRepo *repositories.MealRepository) *MealService {
	return &MealService{
		userRepo: userRepo,
		mealRepo: mealRepo,
	}
}

// Create a new meal
func (s *MealService) CreateMeal(ctx context.Context, userID string, meal *models.Meal, file multipart.FileHeader) error {
	if meal.MealType == "" || meal.MealDate == primitive.DateTime(0) || meal.Items == nil {
		return custom_errors.ErrMissingFields
	}

	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_errors.ErrInvalidObjectID
	}

	meal.CreatedBy = userObjectId

	imageUrl, err := utils.UploadFileToS3(&file)
	if err != nil {
		return err
	}

	meal.PhotoURL = imageUrl

	return s.mealRepo.Create(ctx, meal)
}

// update a meal
func (s *MealService) UpdateMealFromOrder(ctx context.Context, userID string, mealID string, file multipart.FileHeader, mealType string, withPartner string) error {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_errors.ErrInvalidObjectID
	}

	mealObjectId, err := primitive.ObjectIDFromHex(mealID)
	if err != nil {
		return custom_errors.ErrInvalidObjectID
	}

	// check if original meal exists
	originalMeal, err := s.mealRepo.FindByID(ctx, mealObjectId)
	if err != nil {
		return err
	}

	// check if the meal is created by the user
	if originalMeal.CreatedBy != userObjectId {
		return custom_errors.ErrUnauthorized
	}

	imageUrl, err := utils.UploadFileToS3(&file)
	if err != nil {
		return err
	}

	originalMeal.PhotoURL = imageUrl
	originalMeal.MealType = models.MealTpye(mealType)
	originalMeal.WithPartner = withPartner == "true"

	return s.mealRepo.Update(ctx, originalMeal)
}
