package services

import (
	"context"
	"errors"

	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	repo         *repositories.UserRepository
	menuItemRepo *repositories.MenuItemRepository
}

func NewUserService(repo *repositories.UserRepository, menuItemRepo *repositories.MenuItemRepository) *UserService {
	return &UserService{repo: repo, menuItemRepo: menuItemRepo}
}

// Find or create user
func (s *UserService) FindOrCreateUser(ctx context.Context, name, email, googleId, role, picture string) (*models.User, error) {
	user, err := s.repo.FindByGoogleID(ctx, googleId)
	if err == mongo.ErrNoDocuments {
		return s.CreateUser(ctx, name, email, googleId, role, picture)
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

// Create a new user
func (s *UserService) CreateUser(ctx context.Context, name, email, googleID, role, picture string) (*models.User, error) {
	user := models.NewUser(name, email, googleID, role, picture)
	userId, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = userId
	return user, nil
}

// Find a user by ID
func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, objectID)
}

// find a user by email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *UserService) SetChefAndPartner(ctx context.Context, userID string, isChef bool, partnerEmail string) (*models.User, error) {
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.Email == partnerEmail {
		return nil, errors.New("cannot set your partner as yourself")
	}

	// check if parter exists
	partnerUser, err := s.GetUserByEmail(ctx, partnerEmail)
	if err != nil {
		// if partner does not exist, can set isChef to true or false
		if err == mongo.ErrNoDocuments {
			user.IsChef = isChef
		} else {
			return nil, err
		}
	} else {
		// if partner exists, can only set isChef to true if partner is not a chef
		if partnerUser.IsChef && isChef {
			return nil, errors.New("partner is already a chef")
		}

		// if partner is not a chef, user can only be chef
		if !partnerUser.IsChef && !isChef {
			return nil, errors.New("partner is not a chef, you can only be a chef")
		}
		user.IsChef = isChef
	}

	user.PartnerEmail = partnerEmail
	err = s.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
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
