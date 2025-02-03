package services

import (
	"context"

	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo         *repositories.UserRepository
	menuItemRepo *repositories.MenuItemRepository
}

func NewUserService(repo *repositories.UserRepository, menuItemRepo *repositories.MenuItemRepository) *UserService {
	return &UserService{repo: repo, menuItemRepo: menuItemRepo}
}

// Find a user by ID
func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, objectID)
}

func (s *UserService) AddToCart(ctx context.Context, userID, menuItemID string) error {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	menuItemObjectId, err := primitive.ObjectIDFromHex(menuItemID)
	if err != nil {
		return err
	}
	return s.repo.AddToCart(ctx, userObjectId, menuItemObjectId)
}

func (s *UserService) RemoveFromCart(ctx context.Context, userID, menuItemID string) error {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	menuItemObjectId, err := primitive.ObjectIDFromHex(menuItemID)
	if err != nil {
		return err
	}
	return s.repo.RemoveFromCart(ctx, userObjectId, menuItemObjectId)
}

func (s *UserService) GetCartItems(ctx context.Context, userID string) ([]*models.MenuItem, error) {
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	// return empty array if the user just registered
	if len(user.Cart) == 0 {
		return []*models.MenuItem{}, nil
	}

	var createdUserID primitive.ObjectID
	if user.IsChef {
		createdUserID = user.ID
	} else {
		createdUser, err := s.repo.FindByEmail(ctx, user.PartnerEmail)
		if err != nil {
			return nil, err
		}
		createdUserID = createdUser.ID
	}

	var menuItems []*models.MenuItem
	for _, menuItemID := range user.Cart {
		menuItem, err := s.menuItemRepo.GetByID(ctx, menuItemID, createdUserID)
		if err != nil {
			return nil, err
		}
		menuItems = append(menuItems, menuItem)
	}

	return menuItems, nil
}

func (s *UserService) AddOrderToUser(userID, orderID primitive.ObjectID) error {
	return s.repo.AddOrder(userID, orderID)
}

func (s *UserService) RemoveOrderFromUser(userID, orderID primitive.ObjectID) error {
	return s.repo.RemoveOrder(userID, orderID)
}
