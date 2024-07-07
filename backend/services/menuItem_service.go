package services

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	"github.com/peter6866/foodie/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItemService struct {
	repo *repositories.MenuItemRepository
}

func NewMenuItemService(repo *repositories.MenuItemRepository) *MenuItemService {
	return &MenuItemService{repo: repo}
}

// Create a new menu item
func (s *MenuItemService) CreateMenuItem(ctx context.Context, item *models.MenuItem, file multipart.FileHeader) error {
	if item.Name == "" || item.CategoryID.IsZero() {
		return errors.New("name and category ID are required")
	}

	imageUrl, err := utils.UploadFileToS3(&file)
	if err != nil {
		return err
	}

	item.ImageURL = imageUrl

	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now

	return s.repo.Create(ctx, item)
}

func (s *MenuItemService) GetMenuItem(ctx context.Context, id string) (*models.MenuItem, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	return s.repo.GetByID(ctx, objectID)
}

func (s *MenuItemService) GetAllMenuItems(ctx context.Context) ([]*models.MenuItem, error) {
	return s.repo.GetAll(ctx)
}

func (s *MenuItemService) UpdateMenuItem(ctx context.Context, menuItem *models.MenuItem) error {
	if menuItem.ID.IsZero() {
		return errors.New("menu item ID is required")
	}

	menuItem.UpdatedAt = time.Now()

	return s.repo.Update(ctx, menuItem)
}

func (s *MenuItemService) DeleteMenuItem(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID")
	}
	return s.repo.Delete(ctx, objectID)
}

func (s *MenuItemService) GetMenuItemsByCategory(ctx context.Context, categoryID string) ([]*models.MenuItem, error) {
	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	return s.repo.GetByCategory(ctx, objectID)
}

func (s *MenuItemService) GetMenuItemsBySpiceLevel(ctx context.Context, spiceLevel models.SpiceLevel) ([]*models.MenuItem, error) {
	return s.repo.GetBySpiceLevel(ctx, spiceLevel)
}

func (s *MenuItemService) GetMenuItemsByAlcoholContent(ctx context.Context, alcoholContent models.AlcoholContent) ([]*models.MenuItem, error) {
	return s.repo.GetByAlcoholContent(ctx, alcoholContent)
}
