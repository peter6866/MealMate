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
	menuItemRepo *repositories.MenuItemRepository
	categoryRepo *repositories.CategoryRepository
}

func NewMenuItemService(menuItemRepo *repositories.MenuItemRepository, categoryRepo *repositories.CategoryRepository) *MenuItemService {
	return &MenuItemService{
		menuItemRepo: menuItemRepo,
		categoryRepo: categoryRepo,
	}
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

	return s.menuItemRepo.Create(ctx, item)
}

// Custom struct to return menu item with category name
type MenuItemWithCategory struct {
	*models.MenuItem
	CategoryName string `json:"categoryName"`
}

func (s *MenuItemService) GetMenuItem(ctx context.Context, id string) (*MenuItemWithCategory, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	menuItem, err := s.menuItemRepo.GetByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	category, err := s.categoryRepo.GetByID(ctx, menuItem.CategoryID)
	if err != nil {
		return nil, err
	}

	return &MenuItemWithCategory{
		MenuItem:     menuItem,
		CategoryName: category.Name,
	}, nil
}

func (s *MenuItemService) GetAllMenuItems(ctx context.Context) ([]*MenuItemWithCategory, error) {
	menuItems, err := s.menuItemRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	menuItemsWithCategory := make([]*MenuItemWithCategory, len(menuItems))
	for i, item := range menuItems {
		category, err := s.categoryRepo.GetByID(ctx, item.CategoryID)
		if err != nil {
			return nil, err
		}
		menuItemsWithCategory[i] = &MenuItemWithCategory{
			MenuItem:     item,
			CategoryName: category.Name,
		}
	}

	return menuItemsWithCategory, nil
}

func (s *MenuItemService) UpdateMenuItem(ctx context.Context, menuItem *models.MenuItem) error {
	if menuItem.ID.IsZero() {
		return errors.New("menu item ID is required")
	}

	menuItem.UpdatedAt = time.Now()

	return s.menuItemRepo.Update(ctx, menuItem)
}

func (s *MenuItemService) DeleteMenuItem(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID")
	}

	menuItem, err := s.menuItemRepo.GetByID(ctx, objectID)
	if err != nil {
		return err
	}

	if menuItem == nil {
		return errors.New("menu item not found")
	}

	err = utils.DeleteFileFromS3(menuItem.ImageURL)
	if err != nil {
		return err
	}

	return s.menuItemRepo.Delete(ctx, objectID)
}

// TODO: Implement the following methods
func (s *MenuItemService) GetMenuItemsByCategory(ctx context.Context, categoryID string) ([]*models.MenuItem, error) {
	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	return s.menuItemRepo.GetByCategory(ctx, objectID)
}

func (s *MenuItemService) GetMenuItemsBySpiceLevel(ctx context.Context, spiceLevel models.SpiceLevel) ([]*models.MenuItem, error) {
	return s.menuItemRepo.GetBySpiceLevel(ctx, spiceLevel)
}

func (s *MenuItemService) GetMenuItemsByAlcoholContent(ctx context.Context, alcoholContent models.AlcoholContent) ([]*models.MenuItem, error) {
	return s.menuItemRepo.GetByAlcoholContent(ctx, alcoholContent)
}
