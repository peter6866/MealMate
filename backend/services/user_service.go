package services

import (
	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Create a new user
func (s *UserService) CreateUser(name, email, googleID, role, picture string) (*models.User, error) {
	user := models.NewUser(name, email, googleID, role, picture)
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Find a user by ID
func (s *UserService) GetUser(id primitive.ObjectID) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) AddOrderToUser(userID, orderID primitive.ObjectID) error {
	return s.repo.AddOrder(userID, orderID)
}

func (s *UserService) RemoveOrderFromUser(userID, orderID primitive.ObjectID) error {
	return s.repo.RemoveOrder(userID, orderID)
}
