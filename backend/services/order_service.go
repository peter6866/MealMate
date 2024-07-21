package services

import (
	"context"
	"time"

	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	userRepo  *repositories.UserRepository
	orderRepo *repositories.OrderRepository
}

func NewOrderService(userRepo *repositories.UserRepository, orderRepo *repositories.OrderRepository) *OrderService {
	return &OrderService{
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}
}

// Create a new order
func (s *OrderService) CreateOrder(ctx context.Context, userID string, order *models.Order) error {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	order.CreatedBy = userObjectId
	user, err := s.userRepo.FindByID(ctx, userObjectId)
	if err != nil {
		return err
	}

	if !user.IsChef {
		chefUser, err := s.userRepo.FindByEmail(ctx, user.PartnerEmail)
		if err != nil {
			return err
		}
		order.SendTo = chefUser.ID
	} else {
		order.SendTo = user.ID
	}

	order.Status = models.OrderStatusStarted
	order.OrderDate = time.Now()

	return s.orderRepo.Create(ctx, order)
}

// get all orders
func (s *OrderService) GetAllOrders(ctx context.Context, userID string) ([]*models.Order, error) {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(ctx, userObjectId)
	if err != nil {
		return nil, err
	}

	if user.IsChef {
		return s.orderRepo.GetAllForChef(ctx, userObjectId)
	} else {
		return s.orderRepo.GetAllForUser(ctx, userObjectId)
	}
}
